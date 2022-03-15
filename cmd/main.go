package main

import (
	"MAI-Backend/pkg/cache"
	"fmt"
)

func main() {
	lru := cache.NewLRU(100)

	lru.Set("Jesse", "Pinkman")
	lru.Set("Walter", "White")
	lru.Set("Jesse", "James")
	fmt.Println(lru.Get("Jesse"))

	lru.Remove("Walter")
	fmt.Println(lru.Get("Walter"))
}
