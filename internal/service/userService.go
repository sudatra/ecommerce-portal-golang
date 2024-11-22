package service

import (
	"go-ecommerce/internal/domain"
	"go-ecommerce/internal/dto"
)

type UserService struct {
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	return "fake-token", nil
}

func (s UserService) Login(input interface{}) (string, error) {
	return "", nil
}

func (s UserService) FindUserByEmail(email string) (*domain.User, error) {
	return nil, nil
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
