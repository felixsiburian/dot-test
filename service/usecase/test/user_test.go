package test

import (
	"dot-test/service/model"
	"dot-test/service/usecase"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var uuids = uuid.New()

type MockUser struct {
	mock.Mock
}

type MockTools struct {
	mock.Mock
}

func (m *MockUser) Create(user model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUser) UpdateEmail(email, id string) error {
	args := m.Called(email, id)
	return args.Error(0)
}

func (m *MockUser) FindById(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
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
	mockRepo := new(MockUser)
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

func TestUpdateEmail_Success(t *testing.T) {
	mockRepo := new(MockUser)
	usecase := usecase.NewUserUsecase(mockRepo)

	mockRepo.On("UpdateEmail", "test@example.com", "123").Return(nil)

	err := usecase.UpdateEmail("test@example.com", "123")

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "UpdateEmail", "test@example.com", "123")
}

func TestRetrieveById_Success(t *testing.T) {
	mockRepo := new(MockUser)
	usecase := usecase.NewUserUsecase(mockRepo)

	expectedUser := &model.User{ID: uuids, Email: "test@example.com"} // Use uuid.UUID type for ID

	mockRepo.On("FindById", uuids.String()).Return(expectedUser, nil) // Mock expects string input but returns the correct struct

	user, err := usecase.RetrieveById(uuids.String())

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertCalled(t, "FindById", uuids.String())
}
