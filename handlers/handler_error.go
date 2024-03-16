package handlers

import (
	"net/http"

	"github.com/pemba1s1/go-server/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 500, "Something went wrong")
}
