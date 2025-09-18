package main

import (
	"github.com/joho/godotenv"
	"go/adv-example/configs"
	model "go/adv-example/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	config := configs.NewConfig()
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	err = db.AutoMigrate(&model.Link{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
