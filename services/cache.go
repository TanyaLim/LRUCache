package services

import (
	"container/list"
	"fmt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
)

type CacheItem struct {
	n     int
	key   string
	value string
}

type LRUCache struct {
	cache     *list.List
	unique    []string
	cacheSize int
}

func (l *LRUCache) NewLRUCache(cacheSize int) *list.List {
	l.unique = make([]string, 0, 0)
	l.cacheSize = cacheSize

	li := list.New()
	l.cache = li
	return li

}

func (l *LRUCache) PrintCache() {
	fmt.Println("***********")

	for element := l.cache.Front(); element != nil; element = element.Next() {
		tempItem := element.Value.(CacheItem)
		fmt.Printf("Element - %s, key - %s, be found - %d", tempItem.value, tempItem.key, tempItem.n)
		fmt.Println()
	}

	//fmt.Println("***********")
	//for element := l.cache.Front(); element != nil; element = element.Next() {
	//	fmt.Println(element.Value)
	//}
	//fmt.Println()
}

func (l *LRUCache) Add(key string, value string) bool {
	if isUnique(l.unique, key) {
		if l.cache.Len() < l.cacheSize {
			el := l.cache.PushFront(CacheItem{
				n:     1,
				key:   key,
				value: value,
			})
			if el == nil {
				return false
			}
			l.unique = append(l.unique, key)
			return true
		} else {
			el := l.cache.PushFront(CacheItem{
				n:     1,
				key:   key,
				value: value,
			})
			if el == nil {
				return false
			}
			l.unique = append(l.unique, key)

			tempEl := l.cache.Back().Value.(CacheItem)
			l.cache.Remove(l.cache.Back())
			l.unique = remoteUniqueKey(l.unique, tempEl.key)
			return true
		}
	} else {
		for element := l.cache.Front(); element != nil; element = element.Next() {
			tempItem := element.Value.(CacheItem)
			if tempItem.key == key {
				tempItem.n++
				element.Value = tempItem
				for elem := element; elem != nil; elem = elem.Next() {
					if tempItem.n > elem.Value.(CacheItem).n {
						l.cache.MoveAfter(element, elem)
					}
				}
				return true
			}
		}
		return true
	}
}

func isUnique(unique []string, key string) bool {
	for _, u := range unique {
		if u == key {
			return false
		}
	}
	return true
}

func remoteUniqueKey(unique []string, key string) []string {
	fmt.Println("Key: ")
	fmt.Println(key)
	fmt.Printf(InfoColor, "Исходный слайз \n")
	fmt.Println(unique)
	for ixd, u := range unique {
		if u == key {
			unique = append(unique[:ixd], unique[ixd+1:]...)
			fmt.Printf(WarningColor, "После удаления слайз \n")
			fmt.Println(unique)
			fmt.Println()
			return unique
		}
	}
	return unique
}
