package service

import (
	"errors"
	"go-ecommerce/internal/domain"
	"go-ecommerce/internal/dto"
	"go-ecommerce/internal/helper"
	"go-ecommerce/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password);
	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email: input.Email,
		Password: hashedPassword,
		Phone: input.Phone,
	})

	if err != nil {
		return "", errors.New("unable to create user");
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.FindUser(email);
	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.FindUserByEmail(email);

	if err != nil {
		return "", errors.New("user with given email and password does not exist")
	}

	err = s.Auth.VerifyPassword(password, user.Password);
	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input interface{}) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) UpdateProfile(id uint, input interface{}) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input interface{}) (string, error) {
	return "", nil
}

func (s UserService) FindCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input interface{}, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, uid uint) (interface{}, error) {
	return nil, nil
}
