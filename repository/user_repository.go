package repository

import (
	"context"
	"database/sql"

	"github.com/sultansyah/blog-api/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, userUsername string) (domain.User, error)
}
