package main

import (
	"net/http"

	"github.com/Kayky18/API_FULLCYCLE/configs"
	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/database"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8080", nil)

}
