package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/jumagaliev1/one_edu/internal/config"
	"github.com/jumagaliev1/one_edu/internal/logger"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type UserService struct {
	repo   *storage.Storage
	cfg    config.Config
	logger logger.RequestLogger
}

func NewUserService(repo *storage.Storage, cfg config.Config, logger logger.RequestLogger) *UserService {
	return &UserService{repo: repo, cfg: cfg, logger: logger}
}

func (s *UserService) Create(ctx context.Context, user model.User) (*model.User, error) {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	s.logger.Logger(ctx).Info("User password hash", user.Password)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return nil, err
	}
	return s.repo.User.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user model.User) error {
	return s.repo.User.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, ID int) error {
	return s.repo.User.Delete(ctx, ID)
}

func (s *UserService) GetAll(ctx context.Context) ([]*model.User, error) {
	return s.repo.User.GetAll(ctx)
}

func (s *UserService) CheckPassword(encPass, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(providedPassword))
}

func (s *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func (s *UserService) Auth(ctx context.Context, user model.AuthUser) error {
	userFromDB, userErr := s.repo.User.GetByUsername(ctx, user.Username)
	if userErr != nil {
		s.logger.Logger(ctx).Error(userErr)
		return userErr
	}
	checkErr := s.CheckPassword(userFromDB.Password, user.Password)
	if checkErr != nil {
		s.logger.Logger(ctx).Error(checkErr)
		return checkErr
	}

	return nil
}

func (s *UserService) RefreshToken() (string, error) {
	b := make([]byte, 32)

	str := rand.NewSource(time.Now().Unix())
	r := rand.New(str)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *UserService) GenerateToken(user model.AuthUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.JWTClaim{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(s.cfg.JWTKey))
}

func (s *UserService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &model.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.JWTKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*model.JWTClaim)
	if !ok {
		return "", errors.New("token claims are not of type *tokeClaims*")
	}

	return claims.Username, nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.User.GetByUsername(ctx, username)
}

func (s *UserService) GetUserFromRequest(ctx context.Context) (*model.User, error) {
	username, ok := ctx.Value(model.ContextUsername).(string)
	if !ok {
		s.logger.Logger(ctx).Error("not valid context username")
		return nil, errors.New("not valid context username")
	}

	user, err := s.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ChangePassword(ctx context.Context, body model.PasswordReq) error {
	user, err := s.GetUserFromRequest(ctx)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}
	checkErr := s.CheckPassword(user.Password, body.OldPassword)
	if checkErr != nil {
		s.logger.Logger(ctx).Error(checkErr)
		return checkErr
	}

	hash, err := s.HashPassword(body.Password)
	if err != nil {
		s.logger.Logger(ctx).Error(err)
		return err
	}

	user.Password = hash

	return s.Update(ctx, *user)
}
