package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/storage"
	mock_storage "github.com/jumagaliev1/one_edu/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	type fields struct {
		repo *storage.Storage
	}
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				repo: &storage.Storage{User: mock_storage.NewMockIUserRepository(gomock.NewController(t))},
			},
			args:    args{ctx: context.Background(), user: model.User{LastName: "Zhumagaliyev", FirstName: "Alibi", Username: "jumagalibi", Password: "123"}},
			want:    &model.User{LastName: "Zhumagaliyev", FirstName: "Alibi", Username: "jumagalibi", Password: "123"},
			wantErr: false,
		},
		{
			name: "FAIL",
			fields: fields{
				repo: &storage.Storage{User: mock_storage.NewMockIUserRepository(gomock.NewController(t))},
			},
			args:    args{user: model.User{LastName: "Zhumagaliyev", FirstName: "Alibi", Username: "jumagalibi", Password: "123"}},
			want:    &model.User{LastName: "Ruslan", FirstName: "Ruslan", Username: "Ruslan", Password: "123"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMock := mock_storage.NewMockIUserRepository(c)

			newMock.EXPECT().Create(tt.args.ctx, tt.args.user).Return(&tt.args.user, nil)
			got, err := newMock.Create(tt.args.ctx, tt.args.user)
			if err != nil && tt.wantErr {
				return
			}

			assert.Equalf(t, tt.want, got, "Create(%v, %v)", tt.args.ctx, tt.args.user)
		})
	}
}
