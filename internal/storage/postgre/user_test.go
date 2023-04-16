package postgre

import (
	"context"
	"fmt"
	"github.com/jumagaliev1/one_edu/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestUserRepository_GetByID(t *testing.T) {
	db, err := gorm.Open(postgres.Open(fmt.Sprint("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")), &gorm.Config{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&model.User{}, &model.Book{}, &model.Borrow{}, &model.Transaction{})

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ctx context.Context
		ID  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:   "OK",
			fields: fields{DB: db},
			args:   args{ctx: context.Background(), ID: 19},
			want: &model.User{
				ID:        19,
				FirstName: "Alibi",
				LastName:  "Zhumagaliyev",
				Email:     "alibi12@gmail.com",
				Password:  "$2a$14$mVN0Hp5DALtvIQ5zcgd8iOeeFFr6/TAVItqmQegXX70pmTfdtF7Cy",
				Username:  "alibi",
				Balance:   9845000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepository{
				DB: tt.fields.DB,
			}
			got, err := r.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("GetByID() got = %v\nwant %v", got, tt.want)
			}
		})
	}
}
