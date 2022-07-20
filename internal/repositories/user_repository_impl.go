package repositories

import (
	"context"
	"database/sql"
	db "github.com/livoir/goose-migration/db/sqlc"
	"github.com/livoir/goose-migration/internal/core/domain"
	"log"
)

type UserRepositoryImpl struct {
	Queries *db.Queries
}

func NewUserRepositoryImpl(queries *db.Queries) *UserRepositoryImpl {
	return &UserRepositoryImpl{Queries: queries}
}

func (repo *UserRepositoryImpl) Save(ctx context.Context, user domain.User) (domain.User, error) {
	arg := db.CreateUserParams{
		Name: sql.NullString{
			String: user.Name,
			Valid:  true,
		},
		Username: sql.NullString{
			String: user.Username,
			Valid:  true,
		},
		Password: sql.NullString{
			String: user.Password,
			Valid:  true,
		},
	}
	result, err := repo.Queries.CreateUser(ctx, arg)
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}
	user.Id = int(id)
	return user, nil
}

func (repo *UserRepositoryImpl) GetById(ctx context.Context, id int) (domain.User, error) {
	result, err := repo.Queries.GetUserById(ctx, int32(id))

	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	user := domain.User{
		Id:       int(result.ID),
		Username: result.Username.String,
		Name:     result.Name.String,
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetAll(ctx context.Context) ([]domain.User, error) {
	results, err := repo.Queries.GetAllUsers(ctx)

	if err != nil {
		log.Println(err)
		return []domain.User{}, err
	}
	var users []domain.User

	for _, result := range results {
		user := domain.User{
			Id:       int(result.ID),
			Username: result.Username.String,
			Name:     result.Name.String,
			Password: result.Password.String,
		}

		users = append(users, user)
	}

	return users, nil

}

func (repo *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	arg := db.UpdateUserParams{
		Name: sql.NullString{
			String: user.Name,
			Valid:  true,
		},
		ID: int32(user.Id),
	}

	err := repo.Queries.UpdateUser(ctx, arg)

	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, id int) error {
	err := repo.Queries.DeleteUser(ctx, int32(id))

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
