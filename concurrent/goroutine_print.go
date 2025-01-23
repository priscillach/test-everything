package main

import (
	"fmt"
	"sync"
)

func printAnimals() {
	var wg sync.WaitGroup
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch1
			fmt.Println("cat")
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch2
			fmt.Println("dog")
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch3
			fmt.Println("fish")
			if i < 99 {
				ch1 <- struct{}{}
			}
		}
	}()

	// Start the first goroutine
	ch1 <- struct{}{}

	wg.Wait()
}
