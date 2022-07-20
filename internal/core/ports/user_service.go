package ports

import (
	"context"
	"github.com/livoir/goose-migration/internal/core/domain/api"
)

type UserService interface {
	Save(ctx context.Context, request api.UserCreateRequest) (api.UserResponse, error)
	GetById(ctx context.Context, id int) (api.UserResponse, error)
	GetAll(ctx context.Context) ([]api.UserResponse, error)
	Update(ctx context.Context, request api.UserUpdateRequest) (api.UserResponse, error)
	Delete(ctx context.Context, id int) error
}
