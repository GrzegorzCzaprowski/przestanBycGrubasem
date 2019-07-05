package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GrzegorzCzaprowski/przestanBycGrubasem/models"

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

	json.NewDecoder(req.Cookie)
	
	
	
	weight := models.Weight{}
	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print secret message
	fmt.Println("zmiana nastapila")

	// h.M.AddWeight
}
