package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Run1() {
	// 1. dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message.

	// ```go
	// func chanRoutine() {
	// 	log.Print("hello 1")
	// 	go func() {
	// 		time.Sleep(1 * time.Second)
	// 		log.Print("hello 3")
	// 	}
	// 	log.Print("hello 2")
	// }

	// ```

	// -- nâng cao. In ra các message theo thứ tự.
	// -- In ra message 3 trước message 2.
	// Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)
	fmt.Println("\n	1.")

	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
	time.Sleep(2 * time.Second)

	fmt.Println("	In ra message 3 trước message 2: ")
	Cach1()
	Cach2()
	Cach3()
}

func Cach1() {
	fmt.Println("	Cách 1:")
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	time.Sleep(2 * time.Second)
	log.Print("hello 2")
}

func Cach2() {
	fmt.Println("	Cách 2:")
	wg := new(sync.WaitGroup)
	log.Print("hello 1")
	f := func(wg *sync.WaitGroup) {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}
	wg.Add(1)
	go f(wg)
	wg.Wait()
	log.Print("hello 2")
}

func Cach3() {
	fmt.Println("	Cách 3:")
	c1 := make(chan int)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		<-c1
	}()
	c1 <- 1
	log.Println("hello 2")
}
