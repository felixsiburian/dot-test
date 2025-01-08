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

func TestUpdateEmail_Success(t *testing.T) {
	s := setupRedisMock()
	defer s.Close()

	mockRedis := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	mockRepo := new(MockUser)
	usecase := usecase.NewUserUsecase(mockRepo, mockRedis)

	email := "newemail@example.com"
	redisKey := "user-" + uuids.String()
	expectedUser := &model.User{ID: uuids, Email: email}

	mockRedis.Del(redisKey)

	mockRepo.On("UpdateEmail", email, uuids).Return(nil)     // Simulate successful email update
	mockRepo.On("FindById", uuids).Return(expectedUser, nil) // Simulate finding user data

	// Call the function
	err := usecase.UpdateEmail(email, uuids.String())

	// Assertions
	assert.NoError(t, err)
	// Check if the key was set in the Redis store
	storedUser, err := mockRedis.Get(redisKey).Result()
	assert.NoError(t, err)

	// Deserialize and check if the stored data matches
	var user model.User
	err = json.Unmarshal([]byte(storedUser), &user)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Email, user.Email)

	mockRepo.AssertExpectations(t)
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
