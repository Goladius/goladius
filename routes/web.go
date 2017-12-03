package routes

import (
	"net/http"

	"../app/Controllers/Front"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Web(Db *gorm.DB) *mux.Router {

	router := mux.NewRouter()

	// Front //
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Front.HomeController(w, r, Db)
	}).Methods("GET")
	// Documentation

	s := http.StripPrefix("/", http.FileServer(http.Dir("public/")))
	router.PathPrefix("/").Handler(s)
	return router
}
