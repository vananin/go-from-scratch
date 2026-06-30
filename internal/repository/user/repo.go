package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"vananin/go-from-scratch/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		`SELECT id, name, email FROM users WHERE id = $1`,
		id,
	)

	var user domain.User

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}

		return nil, fmt.Errorf("get user %d: %w", id, err)
	}

	return &user, nil
}
