package service

import (
	repository "bookingcontext/store"
	"context"
	"fmt"
	"sync"
)

type BookingService interface {
	StartBooking(ctx context.Context, services []string) (string, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) StartBooking(ctx context.Context, services []string) (string, error) {
	var wg sync.WaitGroup
	result := make(chan string, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, serviceName := range services {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()

			err := s.repo.Book(ctx, name)
			if err != nil {
				fmt.Printf("%s: Booking cancelled.\n", name)
				return
			}

			select {
			case result <- name:
				fmt.Printf("%s: Booking confirmed!\n", name)
				cancel() // 🔥 Cancel other goroutines
			default:
				fmt.Printf("%s: Booking lost race.\n", name)
			}
		}(serviceName)
	}

	select {
	case booked := <-result:
		return booked, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
