package mocks

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	Mock mock.Mock
}

func (pr RepoMock) FindAll() (*models.Users, error) {
	args := pr.Mock.Called()
	return args.Get(0).(*models.Users), args.Error(1)
}

func (pr RepoMock) FindId(id uint64) (*models.User, error) {
	args := pr.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr RepoMock) FindEmail(email string) (*models.User, error) {
	args := pr.Mock.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr RepoMock) Add(data *models.User) (*models.User, error) {
	args := pr.Mock.Called(data)
	return args.Get(0).(*models.User), args.Error(1)
}

func (pr RepoMock) Delete(id uint) error {
	args := pr.Mock.Called(id)
	return args.Error(1)
}

func (pr RepoMock) Update(id uint64, data *models.User) (*models.User, error) {
	args := pr.Mock.Called(id, data)
	return args.Get(0).(*models.User), args.Error(1)
}
