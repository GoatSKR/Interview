package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if rateLimiter.Allow() {
			fmt.Fprintf(w, "Request allowed")
		} else {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
