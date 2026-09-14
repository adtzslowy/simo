package service

import (
	"context"
	"errors"
	"strings"

	"github.com/adtzslowy/simo/internal/auth"
	"github.com/adtzslowy/simo/internal/model"
	"github.com/adtzslowy/simo/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserInactive       = errors.New("invalid credentials")
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtService     *auth.JWTService
}

func NewAuthService(
	userRepository *repository.UserRepository,
	jwtService *auth.JWTService,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *AuthService) Me(
	ctx context.Context,
	userID uuid.UUID,
) (*model.User, error) {
	return s.userRepository.FindByID(ctx, userID)
}

func (s *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.userRepository.FindByEmail(ctx, email)

	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !user.IsActive {
		return "", ErrUserInactive
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(
		user.ID,
		user.Email,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
