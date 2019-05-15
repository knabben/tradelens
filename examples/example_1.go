package main

import (
	"fmt"
	"github.com/knabben/tradelens/token"
	"io/ioutil"
	"log"
	"net/http"
)

var sharedSecret = []byte("test")

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	signature := []byte(r.Header["X-Gtd-Signature"][0])
	isValid := token.VerifySignature(signature, body, sharedSecret)
	fmt.Println("VALID: ", isValid)
	fmt.Fprintf(w, "returning ", isValid)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}