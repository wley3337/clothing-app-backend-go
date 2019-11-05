package app

import (
	"backend/models"
	"context"
	"fmt"
	u "lens/utils"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.request) {
		//end points that do not need authorization
		notAuth := []string{"/api/user/new", "/api/user/login"}
		//current request path
		requestPath := r.URL.Path

		//check if request doesn't need authentication
		for _, value := range notAuth {
			//serve request

			if value == requestPath {
				next.ServeHttp(w, r)
				return
			}
		}

		response := make(map[string]interface{})

		//get the token from the header
		tokenHeader := r.Header.Get("Authorization")

		//if the token is missing: return with a 403 Unauthorized
		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		// token usually comes in 'Bearer {token}'
		splitted := strings.Split(tokenHeader, " ")
		// check to see if the split string is an array of length 2
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//grab the token portion out of the split token header
		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//if there is an error in decoding the token
		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//if the token is invalid, maybe not signed on this server
		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//If Everything went well
		fmt.Sprintf("User %", tk.Username) //<-- Useful for monitoring
		//set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		//proceed in the middleware chain
		next.ServeHttp(w, r)
	})

}
