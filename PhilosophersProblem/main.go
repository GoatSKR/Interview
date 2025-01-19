package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	forks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg.Add(numPhilosophers)
	for _, p := range philosophers {
		go p.dine(ctx, &wg, 500*time.Millisecond)
	}

	wg.Wait()
}
