package config

import (
	"fmt"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/utils"
	"github.com/labstack/gommon/log"
	"github.com/subosito/gotenv"
	"os"
	"time"
)

const (
	defaultHTTPPort = "8000"
)

type ServerConfig struct {
	Host            string
	Port            int
	ShutdownTimeout time.Duration
}
type PostgresConfig struct {
	Username    string
	Password    string
	Host        string
	Port        string
	Database    string
	SSLMode     string
	PingTimeout time.Duration
}
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

func (c PostgresConfig) URI() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SSLMode,
	)
}

func Init(cfg *Config) {

	cfg.Server.Port = utils.StrToInt(os.Getenv("SERVER_PORT"))
	cfg.Server.Host = os.Getenv("SERVER_HOST")
	//cfg.Server.ShutdownTimeout = os.Getenv("SERVER_SHUTDOWN_TIMEOUT")

	cfg.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.Database = os.Getenv("POSTGRES_DATABASE")
	cfg.Postgres.SSLMode = os.Getenv("POSTGRES_SSL_MODE")
}

func New(path string, logger *log.Logger) (*Config, error) {
	cfg := &Config{}

	if err := gotenv.Load(path + "/.env"); err != nil {
		logger.Error(err)
		return nil, err
	}

	Init(cfg)

	return cfg, nil
}
