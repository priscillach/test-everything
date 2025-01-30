package main

import (
	"fmt"
	"sync"
)

func main() {
	deadLock1()
	deadLock2()

	m := sync.Map{}

	m.Range(func(key, value any) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})
}
