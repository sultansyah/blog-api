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

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	sql := "select email, username, password_hash, created_at from users where id = ?"
	rows, err := tx.QueryContext(ctx, sql, userId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			panic(err)
		}

		return user, nil
	}
	return user, errors.New("user not found")
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, userUsername string) (domain.User, error) {
	sql := "select email, username, password_hash, created_at from users where username = ?"
	rows, err := tx.QueryContext(ctx, sql, userUsername)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Username, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			panic(err)
		}

		return user, nil
	}
	return user, errors.New("user not found")
}
