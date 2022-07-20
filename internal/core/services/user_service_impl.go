package services

import (
	"context"
	"github.com/livoir/goose-migration/internal/core/domain"
	"github.com/livoir/goose-migration/internal/core/domain/api"
	"github.com/livoir/goose-migration/internal/core/ports"
	"log"
)

type UserServiceImpl struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Save(ctx context.Context, request api.UserCreateRequest) (api.UserResponse, error) {
	req := domain.User{
		Username: request.Username,
		Name:     request.Name,
		Password: request.Password,
	}

	user, err := service.UserRepository.Save(ctx, req)

	if err != nil {
		log.Println(err)
		return api.UserResponse{}, err
	}

	res := api.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
	}

	return res, nil
}

func (service *UserServiceImpl) GetById(ctx context.Context, id int) (api.UserResponse, error) {
	user, err := service.UserRepository.GetById(ctx, id)

	if err != nil {
		log.Println(err)
		return api.UserResponse{}, err
	}

	res := api.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
	}

	return res, nil
}

func (service *UserServiceImpl) GetAll(ctx context.Context) ([]api.UserResponse, error) {
	users, err := service.UserRepository.GetAll(ctx)

	if err != nil {
		log.Println(err)
		return []api.UserResponse{}, err
	}

	var responses []api.UserResponse
	for _, user := range users {
		res := api.UserResponse{
			Id:       user.Id,
			Name:     user.Name,
			Username: user.Username,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (service *UserServiceImpl) Update(ctx context.Context, request api.UserUpdateRequest) (api.UserResponse, error) {
	req := domain.User{
		Id:   request.Id,
		Name: request.Name,
	}

	user, err := service.UserRepository.Update(ctx, req)

	if err != nil {
		log.Println(err)
		return api.UserResponse{}, nil
	}

	response := api.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
	}

	return response, nil
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) error {
	err := service.UserRepository.Delete(ctx, id)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
