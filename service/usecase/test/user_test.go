package test

import (
	"dot-test/service/model"
	"dot-test/service/usecase"
	"encoding/json"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
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

func (m *MockUser) UpdatePassword(email, id string) error {
	args := m.Called(email, id)
	return args.Error(0)
}

func (m *MockUser) FindById(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUser) Update(user model.User) error {
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

func setupRedisMock() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		log.Fatalf("could not start miniredis server: %v", err)
	}

	return s
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUser)
	mockTools := new(MockTools)

	mockRedisData := setupRedisMock()
	mockRedis := redis.NewClient(&redis.Options{
		Addr: mockRedisData.Addr(),
	})

	usecase := usecase.NewUserUsecase(mockRepo, mockRedis)

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

func TestUpdatePassword_Success(t *testing.T) {
	mockRepo := new(MockUser)
	mockTools := new(MockTools)
	usecase := usecase.NewUserUsecase(mockRepo, nil)

	mockTools.On("HashPassword", "password123").Return("hashedpassword123", nil)
	mockRepo.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil)

	err := usecase.UpdatePassword("hashedpassword123", uuids.String())

	assert.NoError(t, err)
}

func TestRetrieveById_Success(t *testing.T) {
	s := setupRedisMock()
	defer s.Close()

	mockRedis := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	mockRepo := new(MockUser)
	usecase := usecase.NewUserUsecase(mockRepo, mockRedis)

	expectedUser := &model.User{ID: uuids, Email: "test@example.com"} // Use uuid.UUID type for ID

	mockRepo.On("FindById", uuids.String()).Return(expectedUser, nil) // Mock expects string input but returns the correct struct

	user, err := usecase.RetrieveById(uuids.String())

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
}

func TestUpdate_Success(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockUserRepo := new(MockUser)
	mockTools := new(MockTools)

	userUsecase := usecase.NewUserUsecase(mockUserRepo, rdb)

	user := model.User{
		ID:          uuids,
		Email:       "user1@gmail.com",
		Name:        "user 1",
		Phonenumber: "081111111111",
		Username:    "test1",
		Password:    "hashedPass",
	}

	mockTools.On("Validate", user).Return(nil)

	mockUserRepo.On("Update", user).Return(nil)

	updatedUser := &model.User{
		ID:   uuids,
		Name: "user 1 updated",
	}
	mockUserRepo.On("FindById", mock.Anything).Return(updatedUser, nil)

	err = userUsecase.Update(user)

	assert.NoError(t, err)

	redisKey := "user-" + uuids.String()
	data, redisErr := rdb.Get(redisKey).Result()
	assert.NoError(t, redisErr)

	var redisUser model.User
	json.Unmarshal([]byte(data), &redisUser)

	assert.NoError(t, err)
}
