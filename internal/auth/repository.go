package auth

import (
	"errors"
	"github.com/google/uuid"
	"go/adv-example/db"
	"go/adv-example/internal/model"
)

type AuthRepo struct {
	Database *db.Db
}

func NewAuthRepo(database *db.Db) *AuthRepo {
	return &AuthRepo{
		Database: database,
	}
}

func (repo *AuthRepo) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := repo.Database.Find(&user, "email = ?", email)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

func (repo *AuthRepo) GenerateSessionIdByPhone(phone string) (string, error) {
	var user model.User
	result := repo.Database.First(&user, "phone = ?", phone)
	if result.Error != nil {
		return "", result.Error
	}
	sessionId := uuid.New().String()
	user.SessionId = sessionId
	repo.Database.Updates(&user)
	return sessionId, nil
}

func (repo *AuthRepo) FindBySessionId(sessionId, code string) (model.User, error) {
	var user model.User
	result := repo.Database.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		return model.User{}, result.Error
	}

	if code != "1234" {
		return model.User{}, errors.New("invalid code")
	}

	return user, nil
}

func (repo *AuthRepo) Create(user model.User) error {
	result := repo.Database.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
