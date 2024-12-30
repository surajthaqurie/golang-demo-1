package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("mySuperSecrete")

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()

	if err != nil {
		http.Error(w, "Failed to generate JWT: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// If no error, write the valid token to the response writer
	// fmt.Fprintln(w, validToken)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)
	req.Header.Set("Token", validToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error sending request: %s", err.Error())
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(w, "Error: Received status code %d", res.StatusCode)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response body: %s", err.Error())
		return
	}

	fmt.Fprintf(w, "%s", string(body))

}

func GenerateJWT() (string, error) {
	// Create a new token with the specified signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Suraj"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// Generate the signed token string
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		// Return the error if signing fails
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, err
}

func handleRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My simple client")

	handleRequest()
}
