package main

import (
	"database/sql"
	"log"
	"net/http"
	"product-details/handlerlayer"
	"product-details/servicelayer"
	"product-details/storelayer"

	"github.com/gorilla/mux"
)

func main() {

	connstr := "postgres://user1:password1@localhost:5432/user1?sslmode=disable"

	db, err := sql.Open(connstr, "postgres")
	if err != nil {
		log.Println("Database Connection Failed....")
		return
	}
	database := storelayer.NewDatabase(db)
	service := servicelayer.NewService(database)
	handler := handlerlayer.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/products", handler.GetProducts).Methods("GET")
	router.HandleFunc("/products", handler.PostProducts).Methods("POST")
	router.HandleFunc("/products", handler.PutProducts).Methods("PUT")
	router.HandleFunc("/products", handler.DeleteProducts).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
