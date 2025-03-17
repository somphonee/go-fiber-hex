package ports

import (
	"context"

	"github.com/somphonee/go-fiber-hex/internal/core/domain"
)

type UserService interface {
	Create(ctx context.Context, req *domain.CreateUserRequest) (*domain.UserResponse, error)
// 	GetByID(ctx context.Context, id uint) (*domain.UserResponse, error)
// 	Update(ctx context.Context, id uint, req *domain.UpdateUserRequest) (*domain.UserResponse, error)
// 	Delete(ctx context.Context, id uint) error
// 	List(ctx context.Context) ([]*domain.UserResponse, error)
 }
