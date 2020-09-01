package main

import (
	"github.com/faagerholm/fileserver/http"
	"log"
)

const (
	dst = "./media"
)

func main() {
	router := http.SetupRouter()

	err := router.Run(":5000")

	if err != nil {
		log.Fatal(err)
	}
}
