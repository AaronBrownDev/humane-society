package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

// AuthService provides authentication and authorization functionalities, managing user login, registration, and token handling.
type AuthService struct {
	personRepo        domain.PersonRepository
	userRepo          domain.UserAccountRepository
	refreshTokenRepo  domain.RefreshTokenRepository
	roleRepo          domain.RoleRepository
	userRoleRepo      domain.UserRoleRepository
	jwtSecret         []byte
	jwtExpiryMinutes  int
	refreshExpiryDays int
	logger            *log.Logger
}

// NewAuthService initializes and returns a new instance of AuthService with the provided repositories and configuration.
func NewAuthService(personRepo domain.PersonRepository, userRepo domain.UserAccountRepository, refreshTokenRepo domain.RefreshTokenRepository, roleRepo domain.RoleRepository, userRoleRepo domain.UserRoleRepository, jwtSecret []byte, jwtExpiryMinutes int, refreshExpiryDays int) *AuthService {
	return &AuthService{
		personRepo:        personRepo,
		userRepo:          userRepo,
		refreshTokenRepo:  refreshTokenRepo,
		roleRepo:          roleRepo,
		userRoleRepo:      userRoleRepo,
		jwtSecret:         jwtSecret,
		jwtExpiryMinutes:  jwtExpiryMinutes,
		refreshExpiryDays: refreshExpiryDays,
		logger:            log.New(os.Stdout, "auth-service: ", log.LstdFlags),
	}
}

// RegisterRequest represents the input data for user registration
type RegisterRequest struct {
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	EmailAddress    string    `json:"emailAddress"`
	PhysicalAddress string    `json:"physicalAddress"`
	MailingAddress  string    `json:"mailingAddress"`
	BirthDate       time.Time `json:"birthDate"`
	PhoneNumber     string    `json:"phoneNumber"`
	Password        string    `json:"password"`
}

// LoginRequest represents the input data required for a user login, including their email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RefreshTokenRequest represents the input data required for a refresh token request, including the refresh token ID.
type RefreshTokenRequest struct {
	RefreshTokenID string `json:"refreshTokenID"`
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
	// Find the person by email address
	s.logger.Printf("Attempting login for email: %s", req.Email)

	// Use the UserAccountService to get the user account
	userAccountService := NewUserAccountService(s.personRepo, s.userRepo)
	userAccount, err := userAccountService.GetUserAccountByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Printf("Error finding user account for email %s: %v", req.Email, err)
		return nil, errors.New("invalid email")
	}

	if userAccount.IsLocked && userAccount.LockoutEnd.After(time.Now()) {
		return nil, errors.New("account is locked")
	}

	if !userAccount.IsActive {
		return nil, errors.New("account is not active")
	}

	// Log password comparison attempt
	s.logger.Printf("Comparing password hash for user with ID: %s", userAccount.UserID)

	// if the password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(userAccount.PasswordHash), []byte(req.Password)); err != nil {
		s.logger.Printf("Invalid password for user: %s, error: %v", req.Email, err)

		userAccount.FailedLoginAttempts++

		// locks account after 5 failed attempts
		if userAccount.FailedLoginAttempts >= 5 {
			userAccount.IsLocked = true
			userAccount.LockoutEnd = time.Now().Add(time.Minute * 5)
		}

		_, updateErr := s.userRepo.Update(ctx, *userAccount)
		if updateErr != nil {
			s.logger.Printf("Failed to update account after failed login: %v", updateErr)
		}

		return nil, errors.New("invalid password")
	}

	// resets failed login attempts on successful login
	userAccount.FailedLoginAttempts = 0
	userAccount.IsLocked = false
	userAccount.LastLogin = time.Now()
	_, err = s.userRepo.Update(ctx, *userAccount)
	if err != nil {
		s.logger.Printf("Failed to update login timestamp: %v", err)
		// Continue anyway, this is not critical
	}

	s.logger.Printf("Login successful for user: %s", req.Email)

	accessToken, expiresAt, err := s.generateAccessToken(userAccount.UserID)
	if err != nil {
		s.logger.Printf("Error generating access token: %v", err)
		return nil, errors.New("error generating access token")
	}

	refreshToken, err := s.generateRefreshToken(ctx, userAccount.UserID)
	if err != nil {
		s.logger.Printf("Error generating refresh token: %v", err)
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
func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (*domain.UserAccount, error) {
	s.logger.Printf("Registration attempt for email: %s", req.EmailAddress)

	// Check if email already exists
	existingPerson, err := s.personRepo.GetByEmail(ctx, req.EmailAddress)
	if err == nil && existingPerson != nil {
		s.logger.Printf("Email already exists: %s", req.EmailAddress)
		return nil, errors.New("email address already in use")
	} else if err != nil && !errors.Is(err, domain.ErrNotFound) {
		// Only return error if it's not a "not found" error
		s.logger.Printf("Error checking existing email: %v", err)
		return nil, fmt.Errorf("error checking email: %w", err)
	}

	// Create new person record
	personID := uuid.New()
	person := domain.Person{
		PersonID:        personID,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		BirthDate:       req.BirthDate,
		PhysicalAddress: req.PhysicalAddress,
		MailingAddress:  req.MailingAddress,
		EmailAddress:    req.EmailAddress,
		PhoneNumber:     req.PhoneNumber,
	}

	s.logger.Printf("Creating person with ID: %s", personID.String())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Printf("Failed to hash password: %v", err)
		return nil, errors.New("failed to hash password")
	}

	// Create person record
	err = s.personRepo.Create(ctx, &person)
	if err != nil {
		s.logger.Printf("Failed to create person record: %v", err)
		return nil, fmt.Errorf("failed to create person record: %w", err)
	}

	// Create user account with same ID
	userAccount := domain.UserAccount{
		UserID:              personID, // Use the same ID
		PasswordHash:        string(hashedPassword),
		IsActive:            true,
		FailedLoginAttempts: 0,
		IsLocked:            false,
		CreatedAt:           time.Now(),
	}

	s.logger.Printf("Creating user account with ID: %s", personID.String())

	result, err := s.userRepo.Create(ctx, userAccount)
	if err != nil {
		s.logger.Printf("Failed to create user account: %v", err)
		// Try to clean up the orphaned person record
		deleteErr := s.personRepo.Delete(ctx, personID)
		if deleteErr != nil {
			s.logger.Printf("Failed to clean up orphaned person record: %v", deleteErr)
		}
		return nil, fmt.Errorf("failed to create user account: %w", err)
	}

	// Assign "Public" role to the new user
	roleService := NewRoleService(s.roleRepo, s.userRoleRepo)
	err = roleService.AssignRoleToUser(ctx, personID, "Public")
	if err != nil {
		s.logger.Printf("Failed to assign Public role: %v", err)
		// This is not critical to registration, so just log it and continue
	}

	s.logger.Printf("Successfully registered user with email: %s, ID: %s", req.EmailAddress, personID.String())

	return &result, nil
}

// RefreshAccessToken refreshes the access token for the provided refresh token, returning a new AuthResponse if the refresh token is valid.
func (s *AuthService) RefreshAccessToken(ctx context.Context, req RefreshTokenRequest) (*AuthResponse, error) {
	s.logger.Printf("Token refresh attempt for token ID: %s", req.RefreshTokenID)

	refreshTokenID, err := uuid.Parse(req.RefreshTokenID)
	if err != nil {
		s.logger.Printf("Invalid refresh token format: %v", err)
		return nil, errors.New("invalid refresh token format")
	}

	refreshToken, err := s.validateRefreshToken(ctx, refreshTokenID)
	if err != nil {
		s.logger.Printf("Failed to validate refresh token: %v", err)
		return nil, fmt.Errorf("failed to validate refresh token: %w", err)
	}

	accessToken, expiresAt, err := s.generateAccessToken(refreshToken.UserID)
	if err != nil {
		s.logger.Printf("Failed to generate new access token: %v", err)
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	newRefreshToken, err := s.rotateRefreshToken(ctx, refreshToken)
	if err != nil {
		s.logger.Printf("Failed to rotate refresh token: %v", err)
		return nil, fmt.Errorf("failed to rotate refresh token: %w", err)
	}

	s.logger.Printf("Successfully refreshed token for user ID: %s", refreshToken.UserID.String())

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken.TokenID.String(),
		ExpiresAt:    expiresAt,
		UserID:       refreshToken.UserID,
	}, nil
}

// RevokeToken invalidates a refresh token in the database.
// It marks the token as revoked by setting its RevokedAt timestamp to the current time.
// Returns an error if the token doesn't exist or the database operation fails.
func (s *AuthService) RevokeToken(ctx context.Context, tokenID uuid.UUID) error {
	// Find the token
	token, err := s.refreshTokenRepo.GetByTokenID(ctx, tokenID)
	if err != nil {
		return fmt.Errorf("failed to find token: %w", err)
	}

	// If token is already revoked, just return success
	if !token.RevokedAt.IsZero() {
		return nil
	}

	// Mark the token as revoked
	token.RevokedAt = time.Now()

	// Update the token in the database
	_, err = s.refreshTokenRepo.Update(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}

	return nil
}

// Helper methods

// generateAccessToken creates a new JWT access token for the provided userID, returning the token, its expiration, or an error.
func (s *AuthService) generateAccessToken(userID uuid.UUID) (string, time.Time, error) {

	roleService := NewRoleService(s.roleRepo, s.userRoleRepo)
	roleName, err := roleService.GetRoleName(context.Background(), userID)
	if err != nil {
		roleName = "User" // Default role
	}

	expirationTime := time.Now().Add(time.Minute * time.Duration(s.jwtExpiryMinutes))

	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": roleName,
		"exp":  expirationTime.Unix(),
		"iat":  time.Now().Unix(),
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
