package main

import (
	"LRUCache_test/services"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	cacheSize := 8
	rand.Seed(time.Now().UnixNano())

	var c services.LRUCache
	_ = c.NewLRUCache(cacheSize)

	a := 1
	b := 10

	for ixd := 1; ixd < 200; ixd++ {
		i := a + rand.Intn(b-a+1)
		ch := 'a' + rune(rand.Intn('z'-'a'+1))
		c.Add(strconv.Itoa(i), string(ch))
	}
	c.PrintCache()

	fmt.Printf(services.InfoColor, "Result Get(key = 5): ")
	fmt.Println(c.Get("5"))
	fmt.Printf(services.InfoColor, "Result Get(key = 55): ")
	fmt.Println(c.Get("55"))

	fmt.Printf(services.InfoColor, "Result Remove(key = 5): ")
	fmt.Println(c.Remove("5"))
	fmt.Printf(services.InfoColor, "Result Remove(key = 55): ")
	fmt.Println(c.Remove("55"))
	c.PrintCache()
}
