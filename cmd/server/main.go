package main

import (
	"net/http"

	"github.com/Kayky18/API_FULLCYCLE/configs"
	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/database"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperiesIn", configs.JWTExpiresIn))

	r.Route(("/products"), func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))

		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)

		r.Get("/{id}", productHandler.GetProduct)

		r.Get("/", productHandler.GetProducts)

		r.Delete("/{id}", productHandler.DeleteProduct)

		r.Put("/{id}", productHandler.UpdateProduct)
	})

	r.Post("/user", userHandler.CreateUser)
	r.Post("/user/generate-jwt", userHandler.GetJWT)

	http.ListenAndServe(":8080", r)

}
