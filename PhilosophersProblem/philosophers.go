package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

const (
	numPhilosophers = 5
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
}

func (p Philosopher) dine(ctx context.Context, wg *sync.WaitGroup, diningTime time.Duration) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Philosopher %d is done dining\n", p.id)
			return
		default:
			p.think()
			p.eat(diningTime)
		}
	}
}

func (p Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func (p Philosopher) eat(diningTime time.Duration) {
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(diningTime)

	p.rightFork.Unlock()
	p.leftFork.Unlock()
}
