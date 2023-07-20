package services

type ILRUCache interface {
	Add(key string, value string) bool
	PrintCache()
	//	Get(key string) (value string, ok bool)
	//	Remove(key string) (ok bool)
}
