package model

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
)

var (
	ContextUsername = contextKey("username")
)

var ErrDuplicateKey = errors.New("duplicate key not allowed")

type User struct {
	ID        uint           `json:"ID"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Balance   float32        `json:"balance"`
	PhotoURL  string         `json:"photo_URL"`
}

type UserCreateReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type contextKey string

type JWTClaim struct {
	Username       string
	StandardClaims jwt.StandardClaims
}

func (jwt *JWTClaim) Valid() error {
	return nil
}

type PasswordReq struct {
	Password    string `json:"password"`
	OldPassword string `json:"old_password"`
}
