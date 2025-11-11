package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

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
	err := godotenv.Load("./.env")
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

	db.Unscoped().
		Where("id = ?", 1).
		Delete(&model.Order{})
}

func TestCreateOrder(t *testing.T) {
	db := initDb()
	removeData(db)
	CreateUsersInDb(db)
	CreateProductInDb(db)
	ts := httptest.NewServer(App())

	loginPayload, _ := json.Marshal(auth.LoginRequest{
		Email:    "e.konovalov@emcd.io",
		Password: "Google12345^",
	})
	response, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(loginPayload))

	if err != nil {
		t.Fatal(err)
	}
	data, _ := io.ReadAll(response.Body)
	var body auth.LoginResponse
	if err := json.Unmarshal(data, &body); err != nil {
		t.Fatal(err)
	}
	valid, jwtData := jwtPkg.NewJWT(os.Getenv("SECRET")).Parse(body.Token)
	if !valid {
		t.Fatal(errors.New("invalid token"))
	}
	if jwtData.Email != "e.konovalov@emcd.io" {
		t.Fatal(errors.New("invalid token"))
	}

	productsId := []uint{1}
	userId := jwtData.ID
	fmt.Println(userId)
	_, err = CreateWithProducts(db, productsId, userId)

	if err != nil {
		t.Fatal(err)
	}
	defer ts.Close()
}
