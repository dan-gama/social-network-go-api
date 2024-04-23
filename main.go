package main

import (
	"fmt"
	"log"
	"net/http"
	"sn-api/src/config"
	"sn-api/src/router"
)

func main() {
	fmt.Println("API running...")
	config.Init()

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(config.ApiHost, router))
}
