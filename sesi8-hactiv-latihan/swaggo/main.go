package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title Orders API
// @version v0.1
// @description this is simple api test
// @termOfServices http://swagger.io
// @contact.email support@hactiv.com
//
func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
