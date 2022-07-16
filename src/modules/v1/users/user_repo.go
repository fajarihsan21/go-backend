package users

import (
	"errors"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *user_repo {
	return &user_repo{grm}
}

func (r *user_repo) FindAll() (*models.Users, error) {
	var users models.Users
	result := r.db.Order("id_user desc").Find(&users)
	// result := r.db.Find(&users)

	if result.Error != nil {
		return nil, errors.New("error getting all data")
	}
	return &users, nil
}

func (r *user_repo) FindId(id uint64) (*models.User, error) {
	var user models.User
	result := r.db.Where("id_user = ?", id).Find(&user)

	if result.Error != nil {
		return nil, errors.New("failed getting user id")
	}
	return &user, nil
}

func (r *user_repo) FindEmail(email string) (*models.User, error) {
	var users models.User
	result := r.db.First(&users, "email = ?", email)

	if result.Error != nil {
		return nil, errors.New("failed getting email")
	}
	return &users, nil
}

func (r *user_repo) Add(data *models.User) (*models.User, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("error saving the data")
	}
	return data, nil
}

func (r *user_repo) Delete(id uint) error {
	var user models.User
	result := r.db.Where("id_user = ?", id).Delete(&user)
	if result.Error != nil {
		return errors.New("error deleting the data")
	}

	return nil
}

// func (r *user_repo) Update(id uint64, data *models.Users) (*models.Users, error) {
// 	var user models.Users
// 	res := r.db.First(&user, "id_user = ?", id)
// 	if res.Error != nil {
// 		return nil, errors.New("error getting ID")
// 	}

// 	result := r.db.Save(&data)
// 	if result.Error != nil {
// 		return nil, errors.New("error updating data")
// 	}
// 	return data, nil
// }

func (r *user_repo) Update(id uint64, data *models.User) (*models.User, error) {
	var users models.User
	result := r.db.Model(&users).Where("id_user = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("error deleting data")
	}
	return data, nil
}
