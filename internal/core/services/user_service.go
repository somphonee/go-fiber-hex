package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"context"
	"github.com/somphonee/go-fiber-hex/internal/core/domain"

	"github.com/somphonee/go-fiber-hex/internal/core/ports"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{
		userRepo: userRepo,
	}
}
func (s *userService) Create(ctx context.Context, req *domain.CreateUserRequest) error {

	exitstingByEmail, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if exitstingByEmail != nil {
		return errors.New("email already exists")
	}

	existingByUserName, _ := s.userRepo.GetByUserName(ctx, req.Username)
	if existingByUserName != nil {
		return errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return  err
	}

	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return  err
	}
	return nil
}

