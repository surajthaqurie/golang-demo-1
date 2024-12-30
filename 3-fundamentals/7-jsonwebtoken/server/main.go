package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mySuperSecrete")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret information")

}

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				// Ensure the signing method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
				}

				return mySigningKey, nil
			})

			if err != nil {
				// Handle token parsing error
				http.Error(w, "Error parsing token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			if token.Valid {
				// Call the wrapped endpoint if the token is valid
				endpoint(w, r)
			} else {
				// Handle invalid token
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
		} else {
			// Handle missing token
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":9000", nil))

}

func main() {
	fmt.Println("My Simple server")

	handleRequests()
}
