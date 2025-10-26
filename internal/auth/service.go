package auth

import (
	"errors"
	"go/adv-example/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *AuthRepo
}

func NewAuthService(repo *AuthRepo) *AuthService {
	return &AuthService{
		Repo: repo,
	}
}

func (service *AuthService) Register(email, name, password string) (model.User, error) {
	_, err := service.Repo.FindByEmail(email)
	if err != nil {
		return model.User{}, errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}
	userID, err := service.Repo.Create(user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = userID
	return user, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	user, err := service.Repo.FindByEmail(email)
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredentials)
	}

	return user.Email, nil
}

func (service *AuthService) LoginByPhone(phone string) (string, error) {
	sessionId, err := service.Repo.GenerateSessionIdByPhone(phone)
	if err != nil {
		return "", err
	}

	service.Repo.UpdateCode(sessionId)
	return sessionId, nil
}

func (service *AuthService) VerifyPhoneLogin(sessionId, code string) (string, error) {
	user, err := service.Repo.FindBySessionId(sessionId, code)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
