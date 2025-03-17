package ports

import (
	"context"

	"github.com/somphonee/go-fiber-hex/internal/core/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	// GetByID(ctx context.Context, id uint) (*domain.User, error)
  GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetByUserName(ctx context.Context, username string) (*domain.User, error)
	// Update(ctx context.Context, user *domain.User) error
	// Delete(ctx context.Context, id uint) error
	// List(ctx context.Context) ([]*domain.User, error)
}