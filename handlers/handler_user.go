package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pemba1s1/go-server/internal/database"
	"github.com/pemba1s1/go-server/utils"
	"golang.org/x/crypto/bcrypt"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Email == "" || params.UserName == "" || params.Password == "" {
		utils.RespondWithError(w, 400, fmt.Sprintln("Missing Fileds.Couldn't create user."))
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserName:  params.UserName,
		Email:     params.Email,
		Password:  string(password),
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	utils.RespondWithJson(w, 200, user)
}

func (apiCfg *ApiConfig) HandlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.UserName == "" || params.Password == "" {
		utils.RespondWithError(w, 400, fmt.Sprintln("Username or password missing."))
		return
	}

	user, err := apiCfg.DB.GetUserFromUserName(r.Context(), params.UserName)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Couldn't Fetch User: %v", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintln("Username or password is incorrect"))
		return
	}

	utils.RespondWithJson(w, 200, struct {
		UserName string `json:"user_name"`
		Token    string `json:"token"`
	}{
		UserName: user.UserName,
		Token:    "token",
	})
}
