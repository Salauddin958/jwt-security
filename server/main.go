package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":9000", nil)
}

var secretKey = []byte("golangisawesome")

func homepage(w http.ResponseWriter, r *http.Request) {
	if isAuthorized(r, w) {
		fmt.Fprintf(w, "Hello World")
		fmt.Println("Endpoint Hit: homePage")
	} else {
		log.Fatal("Not Authorized ")
	}
}

func isAuthorized(r *http.Request, w http.ResponseWriter) bool {
	if r.Header["Token"] != nil {
		token, err := jwt.Parse(r.Header["Token"][0], func(tkn *jwt.Token) (interface{}, error) {
			if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There is an error in token")
			}
			return secretKey, nil
		})
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		if token.Valid {
			return true
		}
	}
	return false
}
