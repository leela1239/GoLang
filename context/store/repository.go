package repository

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type BookingRepository interface {
	Book(ctx context.Context, serviceName string) error
}

type bookingRepo struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepo{}
}

func (r *bookingRepo) Book(ctx context.Context, serviceName string) error {
	delay := time.Duration(2+rand.Intn(4)) * time.Second
	fmt.Printf("%s: Booking started, will take %v...\n", serviceName, delay)

	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
