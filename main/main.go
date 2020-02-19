package main

import (
	"log"
	"net/http"

	app "starfire/http"
)

func main() {

	prot := &app.WDLoginProtocol{}

	router := app.NewRouter(prot)

	log.Fatal(http.ListenAndServe(":8082", router))
}