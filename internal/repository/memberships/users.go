package memberships

import (
	"context"
	"database/sql"

	"github.com/VH288/miniig/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email string, username string, userID int64) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ? OR id = ?`

	row := r.db.QueryRowContext(ctx, query, email, username, userID)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSERT INTO users (username, email, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.Username, model.Email, model.Password, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
