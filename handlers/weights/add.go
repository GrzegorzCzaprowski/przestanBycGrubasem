package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h WeightHandler) Add(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	session, err := Store.Get(req, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if auth, ok := session.Values["loged"].(bool); !ok || !auth {
		http.Error(w, "You need to log in first!", http.StatusForbidden)
		return
	}

	req.

		// Print secret message
		fmt.Println("zalogowal sie")

	h.M.AddWeight
}
