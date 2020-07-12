package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type Line struct {
	LineNumber int
	Data       string
}

func Run4() {
	// 4. bài tập worker pool: tạo bằng tay file dưới. `file.txt` sau đó đọc từng dòng
	// file này nạp dữ liệu vào 1 buffer channel có size 10, Điều kiện đọc file từng dòng.
	// Chỉ được sử dụng 3 go routine. Kết quả xử lý xong ỉn ra màn hình + từ `xong`
	// ```txt
	// "z9nnHLy8V8"
	// "6AVcSrDUkB"
	// "DezRGPwtx7"
	// "eSmXGjCmTq"
	// "9rfCMntQA5"
	// "Trk6xppMuM"
	// "2sb8BPaUsp"
	// "6AAh6zVFNA"
	// "gsY8kAuKp8"
	// "FQgb8QEpxg"
	// "hEXnKUkYrp"
	// "tchiG2Tiv4"
	// "daMPMJWaM6"
	// "WbBMpX89Sz"
	// "YVnsveajtj"
	// "L9TA7FE5d9"
	// "xBjE7UNe98"
	// "q6bLPeVjYr"
	// "oBppTK62nT"
	// "GxUjEDYBdG"
	// "ZTEpXFStLo"
	// "4XkynbWFvU"
	// "WFmmUSWzDv"
	// "nit8qjmvZH"
	// "iT8BqzHdXo"
	// "7N7mz3qzn2"
	// "KfhMZsHABi"
	// "M4mKWrGgDn"
	// "qLEduDF7so"
	// "YhigrGfLJr"
	// "f82gk2mrxv"
	// "q7TPNZB3Bv"
	// "eWLL5Yg6sG"
	// "GyPqxrXiUg"
	// "86dGJYRzPN"
	// "EWYtAVfXnd"
	// "8dNcD3F8uS"
	// "NLRE6LKqCt"
	// "UbLD2DACiB"
	// "JeLHTTg8vw"
	// ```
	// nâng cao. In ra số lượng goroutine đã khởi tạo.
	// hoàn thiện để tối ưu, thu hồi channel và goroutine đã cấp phát.

	// - Nâng cao 1: Tạo 1 struct `Line` có trường gồm có: `số dòng hiện tại`, `giá trị`
	// của dòng đó.
	// In ra màn hình cú pháp `${line_number} giá trị là: ${data}`.
	// - Nâng cao 2: Khi kết thúc chương trình đã cho đóng những vòng lặp vô hạn của các
	// goroutine lại. Viết chương trình đó.
	// Giợi ý sử dụng biến `make([]chan bool, n)`
	fmt.Println("\n	4.")
	messages := make(chan string, 10)
	wg := &sync.WaitGroup{}
	file, err := os.Open("./file.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	index := 0

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(messages, wg, i, &index)
	}

	for scanner.Scan() {
		messages <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	close(messages)
	wg.Wait()
	fmt.Println("xong")
}

func Worker(messages <-chan string, wg *sync.WaitGroup, i int, index *int) {
	for {
		*index++
		x := Line{*index, ""}
		foo, ok := <-messages
		if !ok {
			*index--
			wg.Done()
			return
		} else {
			x.Data = foo
			println("luồng ", i+1, ":	dòng ", x.LineNumber, " giá trị là: ", x.Data)
		}

	}
}
