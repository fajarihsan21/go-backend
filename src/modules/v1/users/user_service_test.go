package users

import (
	"testing"
	"time"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/fajarihsan21/go-backend/src/modules/v1/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var users models.Users
	repo.Mock.On("FindAll").Return(&users, nil)
	data, err := service.FindAll()

	result := data.Message
	assert.Equal(t, "OK", result, "Expect status = 200")
	assert.Nil(t, err)
}

func TestFindId(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var modelMock = models.User{
		IdUser: 1,
	}

	var expectId uint64 = 1
	repo.Mock.On("FindId", expectId).Return(&modelMock, nil)
	data, err := service.FindId(expectId)

	result := data.Data.(*models.User)
	assert.Equal(t, uint(expectId), result.IdUser, "Expect id_user = 1")
	assert.Nil(t, err)
}

func TestFindEmail(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var modelMock = models.User{
		Email: "user1@email.com",
	}

	repo.Mock.On("FindEmail", "user1@email.com").Return(&modelMock, nil)
	data, err := service.FindEmail("user1@email.com")

	result := data.Data.(*models.User)
	assert.Equal(t, "user1@email.com", result.Email, "Expect email = user1@email.com")
	assert.Nil(t, err)
}

func TestAdd(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var addMock = models.User{
		Email:     "user15@email.com",
		Password:  "1234",
		Name:      "user15",
		Image:     "img15.jpg",
		Birthdate: "15/1/1998",
		Address:   "tangerang",
		Phone:     "98765",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repo.Mock.On("Add", &addMock).Return(&addMock, nil)
	data, err := service.Add(&addMock)

	result := data.Data.(*models.User)
	assert.Equal(t, "user15@email.com", result.Email, "Expect email = user15@email.com")
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var deleteMock = models.User{
		IdUser: 15,
	}

	var deleteId uint64 = 15
	repo.Mock.On("Delete", &deleteMock).Return(&deleteMock, nil)
	data, err := service.Delete(deleteId)

	result := data.Message
	assert.Equal(t, "OK", result, "Expect status = 200")
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	var repo = mocks.RepoMock{Mock: mock.Mock{}}
	var service = user_service{&repo}

	var updateMock = models.User{
		Email:     "user15@email.com",
		Password:  "1234",
		Name:      "user15",
		Image:     "img15.jpg",
		Birthdate: "15/1/1998",
		Address:   "banten",
		Phone:     "98765",
		UpdatedAt: time.Now(),
	}

	var expectId uint64 = 10
	repo.Mock.On("Update", &updateMock).Return(&updateMock, nil)
	data, err := service.Update(expectId, &updateMock)

	result := data.Data.(*models.User)
	assert.Equal(t, "banten", result.Address, "Expect address = banten")
	assert.Nil(t, err)
}
