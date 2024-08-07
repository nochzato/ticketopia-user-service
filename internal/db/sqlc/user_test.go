package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nochzato/ticketopia-user-service/pkg/faker"
	"github.com/nochzato/ticketopia-user-service/pkg/hashpass"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	t.Helper()

	randomUser, err := faker.RandomUser()
	require.NoError(t, err)

	hashed_password, err := hashpass.Hash(randomUser.Password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: randomUser.Username,
		FullName: randomUser.FullName,
		Password: hashed_password,
		Email:    randomUser.Email,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Password, hashed_password)
	require.Equal(t, user.Email, arg.Email)

	require.NotEmpty(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	randomUser := createRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), randomUser.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.ID, randomUser.ID)
	require.Equal(t, user.Username, randomUser.Username)
	require.Equal(t, user.FullName, randomUser.FullName)
	require.Equal(t, user.Password, randomUser.Password)
	require.Equal(t, user.Email, randomUser.Email)
	require.Equal(t, user.CreatedAt, randomUser.CreatedAt)
}

func TestUpdateUserPassword(t *testing.T) {
	randomUser := createRandomUser(t)
	newPassword, err := faker.RandomHashedPassword()
	require.NoError(t, err)

	arg := UpdateUserParams{
		ID: randomUser.ID,
		Password: pgtype.Text{
			String: newPassword,
			Valid:  true,
		},
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEqual(t, user.Password, randomUser.Password)
	require.Equal(t, user.Password, newPassword)
	require.Equal(t, user.ID, randomUser.ID)
	require.Equal(t, user.Username, randomUser.Username)
	require.Equal(t, user.FullName, randomUser.FullName)
	require.Equal(t, user.Email, randomUser.Email)
	require.Equal(t, user.CreatedAt, randomUser.CreatedAt)
}

func TestUpdateUserFullName(t *testing.T) {
	randomUser := createRandomUser(t)
	newFullName := faker.RandomFullName()

	arg := UpdateUserParams{
		ID: randomUser.ID,
		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEqual(t, user.FullName, randomUser.FullName)
	require.Equal(t, user.FullName, newFullName)
	require.Equal(t, user.ID, randomUser.ID)
	require.Equal(t, user.Username, randomUser.Username)
	require.Equal(t, user.Password, randomUser.Password)
	require.Equal(t, user.Email, randomUser.Email)
	require.Equal(t, user.CreatedAt, randomUser.CreatedAt)
}

func TestUpdateUserEmail(t *testing.T) {
	randomUser := createRandomUser(t)
	newEmail := faker.RandomEmail()

	arg := UpdateUserParams{
		ID: randomUser.ID,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEqual(t, user.Email, randomUser.Email)
	require.Equal(t, user.Email, newEmail)
	require.Equal(t, user.ID, randomUser.ID)
	require.Equal(t, user.Username, randomUser.Username)
	require.Equal(t, user.FullName, randomUser.FullName)
	require.Equal(t, user.Password, randomUser.Password)
	require.Equal(t, user.CreatedAt, randomUser.CreatedAt)
}

func TestUpdateAllUserFields(t *testing.T) {
	randomUser := createRandomUser(t)
	newPassword, err := faker.RandomHashedPassword()
	require.NoError(t, err)
	newFullName := faker.RandomFullName()
	newEmail := faker.RandomEmail()

	arg := UpdateUserParams{
		ID:       randomUser.ID,
		FullName: pgtype.Text{String: newFullName, Valid: true},
		Password: pgtype.Text{String: newPassword, Valid: true},
		Email:    pgtype.Text{String: newEmail, Valid: true},
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEqual(t, user.FullName, randomUser.FullName)
	require.NotEqual(t, user.Password, randomUser.Password)
	require.NotEqual(t, user.Email, randomUser.Email)
	require.Equal(t, user.FullName, newFullName)
	require.Equal(t, user.Password, newPassword)
	require.Equal(t, user.Email, newEmail)
	require.Equal(t, user.ID, randomUser.ID)
	require.Equal(t, user.Username, randomUser.Username)
	require.Equal(t, user.CreatedAt, randomUser.CreatedAt)
}
