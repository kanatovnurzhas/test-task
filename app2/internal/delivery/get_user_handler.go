package delivery

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	user, err := h.services.GetUserByEmail(email)
	if err != nil {
		if user == nil {
			message := "user not found"
			messageResponse(w, message, http.StatusNotFound)
			return
		}
		message := "failed get function"
		messageResponse(w, message, 500)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	response := map[string]string{
		"email":    user.Email,
		"salt":     user.Salt,
		"password": user.Password,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
