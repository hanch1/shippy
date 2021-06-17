package main

// 实现数据库的基本 CURD 操作

import (
	"gorm.io/gorm"
	"shippy/user/model"
)

type Repository interface {
	Create(user *model.User) error
	Get(id string) (*model.User, error)
	GetAll() ([]*model.User, error)
	GetByEmailAndPassword(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Create(user *model.User) error {
	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Get(id string) (*model.User, error) {
	var user = &model.User{
		Id: id,
	}
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetByEmailAndPassword(user *model.User) error {
	if err := repo.db.Find(&user).Error; err != nil {
		return err
	}
	return nil
}
