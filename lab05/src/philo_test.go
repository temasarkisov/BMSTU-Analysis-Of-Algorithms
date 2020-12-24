package main

import (
	"testing"
)

func BenchmarkPhilos1(b *testing.B) {
	seatsNum := 10
	eatTimes := 1
	// Initialize chop sticks
	chopSticks := make([]*ChopStick, seatsNum)
	for i := 0; i < seatsNum; i++ {
		chopSticks[i] = new(ChopStick)
	}
	// Initialize chop philosophers
	philos := make([]*Philo, seatsNum)
	for i := 0; i < seatsNum; i++ {
		philos[i] = &Philo{i + 1, chopSticks[i], chopSticks[(i+1)%seatsNum]}
	}

	callHost(philos, seatsNum, eatTimes)
}
