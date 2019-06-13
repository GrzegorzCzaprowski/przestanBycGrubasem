package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
)

var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

func (h UserHandler) Logout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, _ := store.Get(req, "cookie-name")

	// Revoke users authentication
	session.Values["loged"] = false
	session.Save(req, w)
}
