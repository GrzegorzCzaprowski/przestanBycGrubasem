package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Get(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	users := h.M.GetAllUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
}
