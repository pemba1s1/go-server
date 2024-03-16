package handlers

import (
	"net/http"

	"github.com/pemba1s1/go-server/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}
