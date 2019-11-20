package app

import (
	"net/http"
	"context"
	"fmt"
	"../models"
	u "../utils"
	"os"
	"strings"	

	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request){
		noAuth := []string{"/api/driver/new", "/api/driver/login"} //dont auth these endpoints
		requestPath := r.URL.Path

		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")
		
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, "")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-type", "application/json")
			
		}


		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		fmt.Sprintf("Driver %", tk.UserID)
		ctx := context.WithValue(r.Context(), "driver", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
} 