package main

import (
	"fmt"
	"log"
	"sync"
)

func Run3() {
	// 	3. chạy đoạn chương trình dưới đây. Nếu có lỗi hãy thêm logic để nó chạy đúng.
	// - Lý giải nguyên nhân lỗi.

	// ```go
	// func errFunc() {
	// 	m := make(map[int]int)
	// 	for i := 0; i < 1000; i++ {
	// 		go func() {
	// 			for j := 1; j < 10000; j++ {
	// 				if _, ok := m[j]; ok {
	// 					delete(m, j)
	// 					continue
	// 				}
	// 				m[j] = j * 10
	// 			}
	// 		}()
	// 	}

	// 	log.Print("done")
	// }

	// ```

	// -- nâng cao. In ra các message theo thứ tự.
	// -- In ra message 3 trước message 2.
	// Sử dụng 3 cách để làm( gợi ý: sử dụng mutex, chan, waitGroup)
	fmt.Println("\n	3.")
	ErrFunc()
}

func ErrFunc() {
	var mutex = &sync.Mutex{}
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					continue
				}
				m[j] = j * 10
				mutex.Unlock()
			}
		}()
	}

	log.Print("done")
	fmt.Println("độ dài khởi tạo được: ", len(m))
}
