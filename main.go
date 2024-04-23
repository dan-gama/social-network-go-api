package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"sn-api/src/config"
	"sn-api/src/router"
)

func main() {
	fmt.Println("API running...")
	config.Init()

	fmt.Println(generateSecret())

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(config.ApiHost, router))
}

func generateSecret() string {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
