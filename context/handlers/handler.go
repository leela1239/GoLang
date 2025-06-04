package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"bookingcontext/service"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(svc service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: svc}
}

// HTTP handler method
func (h *BookingHandler) HandleBookAny(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle Ctrl+C for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\nUser cancelled booking.")
		cancel()
	}()

	services := []string{"Cab", "Auto", "Bike"}
	booked, err := h.bookingService.StartBooking(ctx, services)
	if err != nil {
		http.Error(w, "Booking failed or cancelled", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Booking successful for: %s\n", booked)
	fmt.Printf("\n%s booking succeeded first. Cancelling others...\n", booked)

	// Optional: allow time for goroutines to clean up
	time.Sleep(1 * time.Second)
	fmt.Println("Booking process complete.")
}
