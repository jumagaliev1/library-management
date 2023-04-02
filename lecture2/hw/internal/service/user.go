package service

import (
	"context"
	"errors"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
)

func (s *Service) GetByID(ctx context.Context, id int) (*model.User, error) {
	usr, err := s.repo.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUserNotFound):
			s.logger.Error(model.ErrUserNotFound)
			return nil, err
		default:
			s.logger.Error(err)
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
			s.logger.Error(model.ErrEmailAlreadyExists)
			return nil, err
		default:
			s.logger.Error(err)
			return nil, err
		}
	}
	return usr, err
}
