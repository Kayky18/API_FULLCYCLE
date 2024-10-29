package main

import (
	"net/http"

	"github.com/Kayky18/API_FULLCYCLE/configs"
	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/database"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	configs, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

	r := chi.NewRouter()

	r.Post("/products", productHandler.CreateProduct)

	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)

	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Put("/products/{id}", productHandler.UpdateProduct)

	r.Post("/user", userHandler.CreateUser)
	r.Post("/user/generate-jwt", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)

}
