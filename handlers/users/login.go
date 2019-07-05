package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"
	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	user := models.User{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Println("error with decoding user to json: ", err)
		w.WriteHeader(500)
		return
	}
	session, err := Store.Get(req, "cookie-name")

	if h.M.LoginUser(user) {
		// Authentication goes here
		// ...

		// Set user as authenticated

		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: user.Email,
		}

		log.Println("password is correct")
		session.Values["loged"] = true
		session.Save(req, w)
		log.Printf("User %v logged succesfully\n", user.Email)
		return
	}
	log.Println("Can't log user, invalid email or password")
	w.WriteHeader(500)
}
