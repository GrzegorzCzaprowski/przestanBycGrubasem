package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h UserHandler) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, _ := store.Get(req, "cookie-name")

	// Authentication goes here
	// ...

	// Set user as authenticated
	session.Values["loged"] = true
	session.Save(req, w)
}
