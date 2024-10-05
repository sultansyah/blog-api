package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sultansyah/blog-api/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := "insert into users (email, username, password_hash) values(?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, user.Email, user.Username, user.PasswordHash)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	sql := "select email, password_hash from users where username = ?"
	rows, err := tx.QueryContext(ctx, sql, user.Username)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	userVerify := domain.User{}
	if rows.Next() {
		err := rows.Scan(&userVerify.Email, &userVerify.PasswordHash)
		if err != nil {
			panic(err)
		}

		if user.PasswordHash == userVerify.PasswordHash {
			return user, nil
		} else {
			return user, errors.New("username and password are incorrect")
		}
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, userUsername string) (domain.User, error) {
	sql := "select password from users where username = ?"
	rows, err := tx.QueryContext(ctx, sql, userUsername)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.CreatedAt)
		if err != nil {
			panic(err)
		}

		return user, nil
	}
	return user, errors.New("user not found")
}
