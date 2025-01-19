package main

import "log"

func main() {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	log.Println(cache.Get(1)) // returns 1

	cache.Put(3, 3)           // evicts key 2
	log.Println(cache.Get(2)) // returns -1 (not found)

	cache.Put(4, 4)           // evicts key 1
	log.Println(cache.Get(1)) // returns -1 (not found)
	log.Println(cache.Get(3)) // returns 3
	log.Println(cache.Get(4)) // returns 4
}
