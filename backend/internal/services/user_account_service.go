package services

import (
	"context"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
)

// UserAccountService provides an abstraction layer for user account operations
// helping to standardize UUID handling between repositories
type UserAccountService struct {
	personRepo      domain.PersonRepository
	userAccountRepo domain.UserAccountRepository
}

// NewUserAccountService creates a new user account service instance
func NewUserAccountService(personRepo domain.PersonRepository, userAccountRepo domain.UserAccountRepository) *UserAccountService {
	return &UserAccountService{
		personRepo:      personRepo,
		userAccountRepo: userAccountRepo,
	}
}

// GetUserAccountByEmail finds a user account via email address
// This provides a reliable way to get a user account during login
func (s *UserAccountService) GetUserAccountByEmail(ctx context.Context, email string) (*domain.UserAccount, error) {
	// Find the person by email
	person, err := s.personRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error finding person by email: %w", err)
	}

	// Log the UUID to help with debugging
	fmt.Printf("Found person with ID: %s\n", person.PersonID.String())

	// Find the user account by person ID
	userAccount, err := s.userAccountRepo.GetByID(ctx, person.PersonID)
	if err != nil {
		return nil, fmt.Errorf("error finding user account: %w", err)
	}

	return userAccount, nil
}
