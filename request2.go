package main

import (
	"fmt"
	"sync"
	"time"
)

func Run2() {
	// 	2. tạo 1 biến X `map[string]string` và 3 goroutine cùng thêm dữ liệu vào X.
	// Mỗi goroutine thêm 1000 key khác nhau. Sao cho quá trình đủ 15 key không mất mát
	// dữ liệu.
	// Lưu ý sử dụng mutex
	fmt.Println("\n	2.")
	var X = make(map[string]string)
	var mutex = &sync.Mutex{}
	var wg = &sync.WaitGroup{}

	f := func(j int) {
		for i := 0; i < 1000; i++ {
			key := time.Now().String()
			val := string(i + 1)
			mutex.Lock()
			X[key] = val
			mutex.Unlock()
			time.Sleep(time.Nanosecond)
		}
		wg.Done()
	}
	for j := 0; j < 3; j++ {
		wg.Add(1)
		go f(j)
	}

	wg.Wait()
	mutex.Lock()
	fmt.Println("Độ dài: ", len(X))
	mutex.Unlock()
}
