package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	h := myhandler{}
	muxmain := initializeRoutes(&h)

	port := 8080
	log.Printf("Starting server on :%d", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), muxmain)
	if err != nil {
		log.Println(err)
	}

}

func initializeRoutes(h *myhandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /pet/{petid}", h.GetPet)
	mux.HandleFunc("POST /pet", h.PostPet)
	mux.HandleFunc("DELETE /pet/{petid}", h.DeletePet)
	mux.HandleFunc("GET /pet", h.GetPetList)
	return mux
}
