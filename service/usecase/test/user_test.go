package test

import (
	"dot-test/service/model"
	"dot-test/service/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserRepo struct {
	mock.Mock
}

type MockTools struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockTools) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockTools) Validate(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepo)
	mockTools := new(MockTools)

	usecase := usecase.NewUserUsecase(mockRepo)

	mockTools.On("HashPassword", "password123").Return("hashedpassword123", nil)
	mockTools.On("Validate", mock.Anything).Return(nil)
	mockRepo.On("Create", mock.Anything).Return(nil)

	request := model.User{
		Email:       "test@gmail.com",
		Name:        "test",
		Phonenumber: "081231231123",
		Username:    "testuser",
		Password:    "password123",
	}

	err := usecase.Create(request)

	assert.NoError(t, err)
}
