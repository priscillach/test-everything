package generic

import "fmt"

type cache[T any] map[string]T

func NewCache[T any]() cache[T] {
	return make(cache[T])
}
func (c cache[T]) Set(k string, v T) {
	c[k] = v
}

func (c cache[T]) Get(k string) (v T, ok bool) {
	v, ok = c[k]
	return
}

func main() {
	c := NewCache[string]()
	b := make(cache[int64])
	b.Set("key2", 42)
	fmt.Println(b.Get("key2"))
	c.Set("key1", "value1")
	value, exists := c.Get("key1")
	if exists {
		println(value) // Output: value1
	} else {
		println("Key not found")
	}
}
