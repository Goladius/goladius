package main

import (
	"log"
	"net/http"

	"./Database"
	"./routes"
)

func main() {

	http.Handle("/", routes.Web(Database.Connection("first")))
	log.Fatal(http.ListenAndServe(":3460", nil))
}
