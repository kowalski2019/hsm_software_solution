package main

import (
	"fmt"
	"log"
	"net/http"
	"ssm/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := ":8008"

	r := mux.NewRouter()
	encr := r.PathPrefix("/api/v1/").Subrouter()
	encr.HandleFunc("/encryption", handler.Encryption).Methods(http.MethodPost)

	fmt.Println("\n\nEncryption API Central started on port " + port + "\n")

	log.Fatal(http.ListenAndServe(port,
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}
