package delivery

import (
	"app2/internal/models"
	"app2/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "http://localhost:8080/generate-salt"

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		message := "failed decode request body"
		messageResponse(w, message, http.StatusInternalServerError)
		return
	}

	err := h.services.CheckUserByEmail(user.Email)
	if err != nil {
		if err == service.ErrIsNotValid {
			message := "email is not valid!"
			messageResponse(w, message, http.StatusBadRequest)
			return
		} else if err == service.ErrAlreadyExists {
			message := "email already exists!"
			messageResponse(w, message, http.StatusBadRequest)
			return
		} else {
			message := "failed mongo check"
			messageResponse(w, message, 500)
		}
	}
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Kasyak gde-to")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Println("Kasyak pri decode")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.services.CreateUser(&user); err != nil {
		message := "failed create user"
		messageResponse(w, message, 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"email":    user.Email,
		"salt":     user.Salt,
		"password": user.Password}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func messageResponse(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := map[string]string{"message": message}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
