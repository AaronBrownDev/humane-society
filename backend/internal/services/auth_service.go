package services

import (
	"context"
	"errors"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	personRepo        domain.PersonRepository
	userRepo          domain.UserAccountRepository
	refreshTokenRepo  domain.RefreshTokenRepository
	jwtSecret         []byte
	jwtExpiryMinutes  int
	refreshExpiryDays int
}

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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

	if userAccount.IsActive {
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

	userAccount.FailedLoginAttempts = 0
	userAccount.IsLocked = false
	userAccount.LockoutEnd = time.Now()
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
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
		UserID:       userAccount.UserID,
	}, nil
}

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

func (s *AuthService) generateRefreshToken(ctx context.Context, userID uuid.UUID) (string, error) {
	tokenID := uuid.New()
	tokenString := uuid.New().String()
	expirationTime := time.Now().Add(time.Hour * 24 * time.Duration(s.refreshExpiryDays))

	refreshToken := domain.RefreshToken{
		TokenID:   tokenID,
		UserID:    userID,
		Token:     tokenString,
		Expires:   expirationTime,
		CreatedAt: time.Now(),
	}

	err := s.refreshTokenRepo.Create(ctx, &refreshToken)
	if err != nil {
		return "", errors.New("failed to save refresh token")
	}

	return tokenString, nil
}
