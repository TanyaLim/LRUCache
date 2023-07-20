package main

import (
	"LRUCache_test/services"
	"math/rand"
	"strconv"
)

func main() {
	cacheSize := 8

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
}
