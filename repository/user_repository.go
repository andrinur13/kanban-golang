package repository

import (
	"kanban-golang/model/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(input entity.User) (entity.User, error)
	CheckSameEmail(email string) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckSameEmail(email string) (entity.User, error) {
	userSame := entity.User{}

	err := r.db.Where("email = ?", email).Find(&userSame).Error

	if err != nil {
		return entity.User{}, err
	}

	return userSame, nil
}

func (r *userRepository) GetByEmail(email string) (entity.User, error) {
	userResult := entity.User{}

	err := r.db.Where("email = ?", email).Find(&userResult).Error

	if err != nil {
		return entity.User{}, err
	}

	return userResult, nil
}
