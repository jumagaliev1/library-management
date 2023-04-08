package postgre

import (
	"context"
	"database/sql"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/config"
	_ "github.com/lib/pq"
	"time"
)
// это должно быть в internal s
func NewClient(cfg config.PostgresConfig) (*sql.DB, error) {
	postgre, err := sql.Open("postgres", cfg.URI())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.PingTimeout+5*time.Second)

	defer cancel()

	err = postgre.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return postgre, nil
}
