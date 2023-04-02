package memory

import (
	"context"
	"github.com/jumagaliev1/one_sdu/lecture2/hw/internal/model"
	"github.com/labstack/gommon/log"
	"sync"
)

type Memory struct {
	users  map[int]*model.User //to do users
	mu     sync.Mutex
	logger *log.Logger
}

func New(logger *log.Logger) *Memory {
	return &Memory{
		users:  make(map[int]*model.User),
		logger: logger,
	}
}

func (r *Memory) Create(ctx context.Context, m map[string]interface{}) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	usr, err := model.NewUser(m)
	if err != nil {
		return nil, err
	}

	for _, u := range r.users {
		if u.Email == usr.Email {
			r.logger.Error(model.ErrEmailAlreadyExists)
			return nil, model.ErrEmailAlreadyExists
		}
	}

	usr.ID = len(r.users) + 1

	r.users[usr.ID] = usr

	return usr, err
}

func (r *Memory) GetByID(ctx context.Context, id int) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	usr, ok := r.users[id]
	if !ok {
		r.logger.Error(model.ErrUserNotFound)
		return nil, model.ErrUserNotFound
	}

	return usr, nil
}
