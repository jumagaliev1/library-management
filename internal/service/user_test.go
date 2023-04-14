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
			args:    args{ctx: context.Background(), user: model.User{LastName: "Zhumagaliyev", FirstName: "Alibi", Username: "jumagalibi", Password: "$2a$14$O74EkRPilseWCRnvBJnv2uXDL54UtL.EXNRwx2YJ5lRrJTk5eg8tS"}},
			want:    &model.User{LastName: "Zhumagaliyev", FirstName: "Alibi", Username: "jumagalibi", Password: "$2a$14$O74EkRPilseWCRnvBJnv2uXDL54UtL.EXNRwx2YJ5lRrJTk5eg8tS"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			newMock := mock_storage.NewMockIUserRepository(c)
			s := &UserService{
				repo: &storage.Storage{
					User: newMock,
				},
			}
			got, err := s.Create(tt.args.ctx, tt.args.user)
			if err != nil && tt.wantErr {
				return
			}

			newMock.EXPECT().Create(tt.args.ctx, tt.args.user).Return(tt.want, nil)
			//got, err := s.Create(tt.args.ctx, tt.args.user)
			//if err != nil && tt.wantErr {
			//	return
			//}
			assert.Equalf(t, tt.want, got, "Create(%v, %v)", tt.args.ctx, tt.args.user)

		})
	}
}
