package repository

import (
	"errors"
	"go-ecommerce/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	err := r.db.Create(&u).Error
	
	if err != nil {
		log.Printf("Error Creating user %v", err);
		return domain.User{}, errors.New("unable to create user")
	}

	return u, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("Error Finding user %v", err);
		return domain.User{}, errors.New("user does not exist")
	}

	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf("Error Finding user %v", err);
		return domain.User{}, errors.New("user with id does not exist")
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error

	if err != nil {
		log.Printf("Error Updating user %v", err);
		return domain.User{}, errors.New("unable to update user")
	}

	return user, nil
}