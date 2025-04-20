package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthService provides authentication and authorization functionalities, managing user login, registration, and token handling.
type AuthService struct {
	personRepo        domain.PersonRepository
	userRepo          domain.UserAccountRepository
	refreshTokenRepo  domain.RefreshTokenRepository
	jwtSecret         []byte
	jwtExpiryMinutes  int
	refreshExpiryDays int
}

// NewAuthService initializes and returns a new instance of AuthService with the provided repositories and configuration.
func NewAuthService(personRepo domain.PersonRepository, userRepo domain.UserAccountRepository, refreshTokenRepo domain.RefreshTokenRepository, jwtSecret []byte, jwtExpiryMinutes int, refreshExpiryDays int) *AuthService {
	return &AuthService{
		personRepo:        personRepo,
		userRepo:          userRepo,
		refreshTokenRepo:  refreshTokenRepo,
		jwtSecret:         jwtSecret,
		jwtExpiryMinutes:  jwtExpiryMinutes,
		refreshExpiryDays: refreshExpiryDays,
	}
}

// LoginRequest represents the input data required for a user login, including their email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse represents the response returned after a successful authentication or token refresh.
// It includes an access token, refresh token, expiration time, and the associated user ID.
type AuthResponse struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	UserID       uuid.UUID `json:"userID"`
}

// Login authenticates the user's credentials and returns an AuthResponse if the credentials are valid.
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {

	// TODO could avoid having a person repo here if UserAccount has email column might be able to implement this with a join instead in the UserAccount repository.
	// TODO might want to expand on the error messages here to be more specific.
	person, err := s.personRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email")
	}

	userAccount, err := s.userRepo.GetByID(ctx, person.PersonID)
	if err != nil {
		return nil, errors.New("no user account found")
	}

	if userAccount.IsLocked && userAccount.LockoutEnd.After(time.Now()) {
		return nil, errors.New("account is locked")
	}

	if !userAccount.IsActive {
		return nil, errors.New("account is not active")
	}

	// if the password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordHash), []byte(req.Password)); err != nil {
		userAccount.FailedLoginAttempts++

		// locks account after 5 failed attempts
		if userAccount.FailedLoginAttempts >= 5 {
			userAccount.IsLocked = true
			userAccount.LockoutEnd = time.Now().Add(time.Minute * 5)
		}

		s.userRepo.Update(ctx, *userAccount)

		return nil, errors.New("invalid password")
	}

	// resets failed login attempts on successful login
	userAccount.FailedLoginAttempts = 0
	userAccount.IsLocked = false
	userAccount.LastLogin = time.Now()
	s.userRepo.Update(ctx, *userAccount)

	accessToken, expiresAt, err := s.generateAccessToken(userAccount.UserID)
	if err != nil {
		return nil, errors.New("error generating access token")
	}

	refreshToken, err := s.generateRefreshToken(ctx, userAccount.UserID)
	if err != nil {
		return nil, errors.New("error generating refresh token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.TokenID.String(),
		ExpiresAt:    expiresAt,
		UserID:       userAccount.UserID,
	}, nil
}

// Register creates a new UserAccount based on the provided Person and password, and persists it in the repositories.
func (s *AuthService) Register(ctx context.Context, person domain.Person, password string) (*domain.UserAccount, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// TODO might want to implement create func for person and userAccount.
	// TODO might create orphan Person if create UserAccount fails. Need to eventually handle that edge case.
	err = s.personRepo.Create(ctx, &person)
	if err != nil {
		return nil, errors.New("failed to create underlying person constraint")
	}

	userAccount := domain.UserAccount{
		UserID:              person.PersonID,
		PasswordHash:        string(hashedPassword),
		IsActive:            true,
		FailedLoginAttempts: 0,
		IsLocked:            false,
		CreatedAt:           time.Now(),
	}

	// TODO have to look into if create should return UserAccount because I didn't think about that initially.
	result, err := s.userRepo.Create(ctx, userAccount)
	if err != nil {
		return nil, errors.New("failed to create user account")
	}

	// TODO this could return userAccount if plan to change create func
	return &result, nil
}

// RefreshAccessToken refreshes the access token for the provided refresh token, returning a new AuthResponse if the refresh token is valid.
func (s *AuthService) RefreshAccessToken(ctx context.Context, refreshTokenIDStr string) (*AuthResponse, error) {
	refreshTokenID, err := uuid.Parse(refreshTokenIDStr)
	if err != nil {
		return nil, errors.New("invalid refresh token format")
	}

	refreshToken, err := s.validateRefreshToken(ctx, refreshTokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to validate refresh token: %w", err)
	}

	accessToken, expiresAt, err := s.generateAccessToken(refreshToken.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	newRefreshToken, err := s.rotateRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to rotate refresh token: %w", err)
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.TokenID.String(),
		ExpiresAt:    expiresAt,
		UserID:       refreshToken.UserID,
	}, nil
}

// Helper methods

// generateAccessToken creates a new JWT access token for the provided userID, returning the token, its expiration, or an error.
func (s *AuthService) generateAccessToken(userID uuid.UUID) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(s.jwtExpiryMinutes))

	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)

	return tokenString, expirationTime, err
}

// generateRefreshToken creates a new JWT refresh token for the provided userID, returning the token or an error.
func (s *AuthService) generateRefreshToken(ctx context.Context, userID uuid.UUID) (*domain.RefreshToken, error) {
	tokenID := uuid.New()
	expirationTime := time.Now().Add(time.Hour * 24 * time.Duration(s.refreshExpiryDays))

	refreshToken := domain.RefreshToken{
		TokenID:   tokenID,
		UserID:    userID,
		Expires:   expirationTime,
		CreatedAt: time.Now(),
	}

	err := s.refreshTokenRepo.Create(ctx, &refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}

	return &refreshToken, nil
}

// validateRefreshToken validates the given refresh token by its ID, checking if it exists, is not expired, and not revoked.
func (s *AuthService) validateRefreshToken(ctx context.Context, tokenID uuid.UUID) (*domain.RefreshToken, error) {
	token, err := s.refreshTokenRepo.GetByTokenID(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// if expired throw error
	if token.Expires.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	// if revoked throw error
	if !token.RevokedAt.IsZero() {
		return nil, errors.New("refresh token revoked")
	}

	return token, nil
}

// rotateRefreshToken revokes the old refresh token and generates a new one, linking them for traceability.
func (s *AuthService) rotateRefreshToken(ctx context.Context, oldToken *domain.RefreshToken) (*domain.RefreshToken, error) {
	newToken, err := s.generateRefreshToken(ctx, oldToken.UserID)
	if err != nil {
		return nil, err
	}

	oldToken.RevokedAt = time.Now()
	oldToken.ReplacedByTokenID = newToken.TokenID

	_, err = s.refreshTokenRepo.Update(ctx, oldToken)
	if err != nil {
		return nil, err
	}

	return newToken, nil
}
