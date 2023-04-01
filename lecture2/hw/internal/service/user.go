package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
)

//type IUserRepository interface {
//	Create(ctx context.Context, m map[string]interface{}) (*model.User, error)
//	GetByID(ctx context.Context, id int) (*model.User, error)
//}
//
//type UserService struct {
//	repo IUserRepository
//}

func (s *Service) GetByID(ctx context.Context, id int) (*model.User, error) {
	usr, err := s.repo.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			//logger
			return nil, err
		default:
			//logger
			return nil, err
		}
	}
	return usr, nil
}

func (s *Service) Create(ctx context.Context, input model.UserInput) (*model.User, error) {
	m := model.SetUser(input.FirstName, input.LastName, input.Email, input.Password)
	usr, err := s.repo.Create(ctx, m)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrEmailAlreadyExists):
			// TO-DO logger
			return nil, err
		default:
			// TO-DO logger
			return nil, err
		}
	}
	return usr, err
}
