package middleware

import (
	"encoding/json"
	"latihan-sesi9/models"
	"latihan-sesi9/params"
	"net/http"
)

func Auth(rw http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	isValid := false

	if !ok {
		outputJson(rw, "something went wrong")
		return false
	}

	userData := models.GetUsers()
	for _, u := range *userData {
		if (u.Username == username) && (u.Password == password) {
			isValid = true
			break
		}
	}

	if !isValid {
		outputJson(rw, "something went wrong")
		return false
	}

	return true
}

func outputJson(rw http.ResponseWriter, payload interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(params.Response{
		Status:  http.StatusUnauthorized,
		Message: "UNAUTHORIZED",
		Payload: payload,
	})
}
