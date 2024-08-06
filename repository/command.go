package repository

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/google/uuid"
)

func (repo *Repository) AddCommand(command *models.Command, post_id string, user_id string) error {
	parsed_post_id, err := uuid.Parse(post_id)
	if err != nil {
		return errors.New("invalid id")
	}
	parsed_user_id, err := uuid.Parse(user_id)
	if err != nil {
		return errors.New("invalid user id")
	}
	command.Post_id = parsed_post_id
	command.User_id = parsed_user_id
	if err := repo.DB.Create(&command).Error; err != nil {
		return errors.New("command creation failed")
	}
	return nil
}

func (repo *Repository) DeleteCommand(id string) error {
	if err := repo.DB.Delete(&models.Command{}, "command_id=?", id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetCommandsByPostId(id string) ([]models.Command, error) {
	var commands []models.Command
	if err := repo.DB.Where("post_id=?", id).Find(&commands).Error; err != nil {
		return nil, err
	}
	return commands, nil
}
