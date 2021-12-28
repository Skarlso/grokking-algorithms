package chapter05

import "fmt"

// Data defines a cacheable data entry.
type Data interface {
	string | []byte
}

// DataCache holds some data in a cache based on an arbitrary string.
type DataCache[T Data] map[string]T

// NewCache creates a new data cache.
func NewCache[T Data]() DataCache[T] {
	return make(DataCache[T])
}

// GetPage uses the data cache to retrieve a cached page.
func GetPage[T Data](dataCache DataCache[T], url string) T {
	v, ok := dataCache[url]
	if ok {
		return v
	}
	data := GetDataFromServer[T](url)
	dataCache[url] = data
	return data
}

// GetDataFromServer is a mock which is called if the cache doesn't have an entry for
// a key.
func GetDataFromServer[T Data](key string) T {
	fmt.Println("getting data for key: ", key)
	return T([]byte("this is generated data"))
}
