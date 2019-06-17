package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user from json: ", err)
		w.WriteHeader(500)
		return
	}
	user = h.M.PostUser(user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("error with encoding to json: ", err)
		w.WriteHeader(500)
		return
	}
}
