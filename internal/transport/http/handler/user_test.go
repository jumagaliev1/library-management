package handler

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jumagaliev1/one_edu/internal/model"
	"github.com/jumagaliev1/one_edu/internal/service"
	mock_service "github.com/jumagaliev1/one_edu/internal/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_Create(t *testing.T) {
	type mockBehavior func(r *mock_service.MockIUserService, user model.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            model.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"first_name": "Alibi", "last_name": "Zhumagaliyev", "username": "jumagalibi", "password": "123"}`,
			inputUser: model.User{
				FirstName: "Alibi",
				LastName:  "Zhumagaliyev",
				Username:  "jumagalibi",
				Password:  "123",
			},
			mockBehavior: func(r *mock_service.MockIUserService, user model.User) {
				r.EXPECT().Create(context.Background(), user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockIUserService(c)
			test.mockBehavior(repo, test.inputUser)

			service := &service.Service{User: repo}
			handler := UserHandler{service: service}

			e := echo.New()
			e.POST("/user", handler.Create)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user",
				bytes.NewBufferString(test.inputBody))

			e.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
	//ID        uint           `json:"ID"`
	//CreatedAt time.Time      `json:"-"`
	//UpdatedAt time.Time      `json:"-"`
	//DeletedAt gorm.DeletedAt `json:"-"`
	//FirstName string         `json:"first_name"`
	//LastName  string         `json:"last_name"`
	//Username  string         `json:"username"`
	//Email     string         `json:"email"`
	//Password  string         `json:"password"`
	//Balance   float32        `json:"balance"`
	//PhotoURL  string         `json:"photo_URL"`
	
}
