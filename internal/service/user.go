package service

import (
	"errors"
	"github.com/AliasgharHeidari/test-gorm/internal/database"
	"github.com/AliasgharHeidari/test-gorm/internal/model"
	"github.com/AliasgharHeidari/test-gorm/internal/security"
)

func CreateUser(inputUser model.User) error {
	bytes, err := security.HashPassword(inputUser.Password)
	if errors.Is(err, security.ErrEmptyPassword) {

		return security.ErrEmptyPassword

	} else if errors.Is(err, security.ErrLengthPassword) {

		return security.ErrLengthPassword

	} else if err != nil {

		return err
	}

	user := model.User{
		UserName: inputUser.UserName,
		Password: bytes,
		Email:    inputUser.Email,
	}

	DB := database.GetDB()
	if err := DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
