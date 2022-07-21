package repositories

import (
	"context"
	"database/sql"
	"github.com/livoir/goose-migration/internal/core/domain"
	"github.com/livoir/goose-migration/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) domain.User {
	name := util.RandomString()
	username := util.RandomString()
	password := util.RandomString()

	user := domain.User{
		Username: username,
		Name:     name,
		Password: password,
	}

	result, err := userRepository.Save(context.Background(), user)

	require.NoError(t, err)
	require.Empty(t, result)
	require.NotZero(t, result.Id)

	return domain.User{
		Id:       result.Id,
		Username: username,
		Name:     name,
		Password: password,
	}
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	expected := createRandomUser(t)

	result, err := userRepository.GetById(context.Background(), expected.Id)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, expected.Id, result.Id)
	require.Equal(t, expected.Name, result.Name)
	require.Equal(t, expected.Username, result.Username)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	newName := util.RandomString()

	newUser := domain.User{
		Name: newName,
		Id:   user.Id,
	}
	res, err := userRepository.Update(context.Background(), newUser)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, newUser.Name, res.Name)

	result, err := userRepository.GetById(context.Background(), user.Id)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, newUser.Name, result.Name)

}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := userRepository.Delete(context.Background(), user.Id)
	require.NoError(t, err)

	result, err := userRepository.GetById(context.Background(), user.Id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, result)

}

func TestGetAllUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	users, err := userRepository.GetAll(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users), 10)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
