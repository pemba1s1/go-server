package handlers

import (
	"net/http"

	"github.com/pemba1s1/go-server/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) middlewareAuth()
