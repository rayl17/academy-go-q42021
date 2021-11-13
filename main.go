package main

import (
	"log"

	"github.com/raymundo/academy-go-q42021/router"
)

func main() {

	srv := router.NewRouter()

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
