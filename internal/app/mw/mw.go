package mw

import (
	"log"
	"net/http"
)

func Role_chek(next http.HandlerFunc) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		user_role := r.Header.Get("User-role")
		if user_role == "admin" {
			log.Print("red buttom")
		}
		next.ServeHTTP(w, r)
	}
}