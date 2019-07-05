package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, _ := Store.Get(req, "cookie-name")

	// Revoke users authentication
	session.Values["loged"] = false
	session.Save(req, w)
	log.Println("user logged out")
}
