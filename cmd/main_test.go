package main

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/joho/godotenv"
	"go/adv-example/internal/auth"
	"go/adv-example/internal/model"
	jwtPkg "go/adv-example/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env.development")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	err = db.AutoMigrate(&model.User{}, &model.Product{}, model.Stat{}, &model.Link{}, &model.Order{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return db
}

func CreateUsersInDb(db *gorm.DB) {
	password := "Google12345^"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	db.Create(&model.User{
		Email:    "e.konovalov@emcd.io",
		Password: string(hashedPassword),
		Name:     "Ed",
	})
}

func CreateProductInDb(db *gorm.DB) {
	db.Create(&model.Product{
		Name:        "iPhone 17 Pro Max",
		Price:       110000,
		Description: "Phone",
	})
}

func CreateWithProducts(db *gorm.DB, productsId []uint, userID uint) (model.Order, error) {
	var products []model.Product
	if result := db.Where("id in ?", productsId).Find(&products); result.Error != nil {
		return model.Order{}, result.Error
	}

	if len(products) != len(productsId) {
		return model.Order{}, errors.New("products count not match")
	}

	order := model.Order{
		UserID:   userID,
		Products: products,
	}

	if result := db.Create(&order); result.Error != nil {
		return model.Order{}, result.Error
	}

	return order, nil
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "e.konovalov@emcd.io").
		Delete(&model.User{})
}

func TestCreateOrder(t *testing.T) {
	db := initDb()
	CreateUsersInDb(db)
	CreateProductInDb(db)

	loginPayload, _ := json.Marshal(auth.LoginRequest{
		Email:    "e.konovalov@emcd.io",
		Password: "Google12345^",
	})
	userJwt, _ := http.Post("http://localhost:8080/create-order", "application/json", bytes.NewReader(loginPayload))
	valid, jwtData := jwtPkg.NewJWT(os.Getenv("SECRET")).Parse(userJwt)
	productsId := []uint{1}
	_, err := CreateWithProducts(db, productsId, 1)
	if err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(App())
	defer ts.Close()

	removeData(db)
}
