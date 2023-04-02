package postgre

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
	"github.com/labstack/gommon/log"
)

var (
	ErrPQDuplicateEmail = `pq: duplicate key value violates unique constraint "users_email_key"`
)

type IPostgresqlClient interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRow(query string, args ...any) *sql.Row
}

type Repository struct {
	client IPostgresqlClient
	logger *log.Logger
}

func New(client IPostgresqlClient, logger *log.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}

func (r *Repository) Create(ctx context.Context, m map[string]interface{}) (*model.User, error) {
	query := `
			INSERT INTO users (first_name, last_name, email, password)
			VALUES ($1, $2, $3, $4)
			RETURNING id`

	args := []interface{}{
		m[model.FirstNameField],
		m[model.LastNameField],
		m[model.EmailField],
		m[model.PasswordField],
	}
	var id int
	err := r.client.QueryRowContext(ctx, query, args...).Scan(&id)

	m[model.IDField] = id
	if err != nil {
		switch {
		case err.Error() == ErrPQDuplicateEmail:
			r.logger.Error(ErrPQDuplicateEmail)
			return nil, model.ErrEmailAlreadyExists
		default:
			r.logger.Error(err)
			return nil, err
		}
	}

	usr, err := model.NewUser(m)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*model.User, error) {
	query := `
			SELECT id, first_name, last_name, email, password
			FROM users
			WHERE id = $1`

	var usr model.User
	err := r.client.QueryRowContext(ctx, query, id).Scan(
		&usr.ID,
		&usr.FirstName,
		&usr.LastName,
		&usr.Email,
		&usr.Password,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			r.logger.Error(sql.ErrNoRows)
			return nil, model.ErrUserNotFound

		default:
			r.logger.Error(err)
			return nil, err
		}
	}

	return &usr, nil

}
