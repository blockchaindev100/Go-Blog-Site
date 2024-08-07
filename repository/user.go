package repository

import (
	"time"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/blockchaindev100/Go-Blog-Site/service"
)

type User interface {
	CreateUser(*models.User) error
	GetUserByEmail(string) (models.User, error)
	GetUserById(string) (models.User, error)
}

func (repo *Repository) GetUserById(id string) (models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, "user_id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *Repository) CreateUser(user *models.User) error {
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		repo.Logger.Error(err)
		return err
	}
	user.Password = hashedPassword
	user.Created_at = time.Now()
	if err := repo.DB.Create(user).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := repo.DB.Where("email=?", email).First(&user).Error; err != nil {
		repo.Logger.Error(err)
		return user, err
	}
	return user, nil
}