package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Kayky18/API_FULLCYCLE/internal/dto"
	"github.com/Kayky18/API_FULLCYCLE/internal/entity"
	"github.com/Kayky18/API_FULLCYCLE/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(u database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: u}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)

}
