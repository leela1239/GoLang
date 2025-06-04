package main

import (
	handler "bookingcontext/handlers"
	"bookingcontext/service"
	repository "bookingcontext/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewBookingRepository()
	service := service.NewBookingService(repo)
	h := handler.NewBookingHandler(service)

	http.HandleFunc("/book-any", h.HandleBookAny)

	fmt.Println(" Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
