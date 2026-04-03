/*
Thử thách 2: Xử lý đồng thời (Concurrency với Goroutines & Channels)
Đây là "vũ khí hạng nặng" của Go khi cần tăng tốc độ xử lý các tác vụ như gọi database hoặc xuất file hàng loạt.

Yêu cầu:

Tạo một channel kiểu int (ch := make(chan int)).

Viết một hàm generateData(n int, ch chan<- int) dùng vòng lặp để đẩy các số từ 1 đến n
(giả lập ID của các bản ghi) vào channel, sau đó close(ch).
Chạy hàm này dưới dạng một Goroutine.

Viết một hàm worker(ch <-chan int, wg *sync.WaitGroup)
để đọc dữ liệu từ channel và in ra màn hình: "Worker đang xử lý bản ghi ID: ...".

Ở hàm main, khởi tạo sync.WaitGroup, tạo 3 workers (chạy đồng thời bằng Goroutines)
để cùng đọc và xử lý 10 bản ghi từ channel trên. Đảm bảo chương trình chờ tất cả
workers chạy xong mới kết thúc.
*/
package main

import (
	"fmt"
	"sync"
)

func generateData(n int, ch chan<- int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

// Thêm tham số workerID để dễ phân biệt ai đang làm việc
func worker(workerID int, ch <-chan int, wg *sync.WaitGroup) {
	// Bắt buộc: Báo cho WaitGroup biết worker này đã làm xong khi hàm kết thúc
	defer wg.Done()

	// Vòng lặp range sẽ tự động thoát khi kênh 'ch' bị đóng
	for id := range ch {
		fmt.Printf("Worker %d đang xử lý bản ghi ID: %d\n", workerID, id)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Khởi chạy bộ phát dữ liệu
	go generateData(10, ch)

	// Tạo 3 workers (Pattern Fan-Out)
	numWorkers := 3
	wg.Add(numWorkers) // Khai báo số lượng worker cần chờ

	for i := 1; i <= numWorkers; i++ {
		// Khởi chạy các Goroutine worker đồng thời
		go worker(i, ch, &wg)
	}

	// Bắt buộc: Chặn hàm main lại, chờ đến khi WaitGroup đếm lùi về 0
	wg.Wait()

	fmt.Println("Toàn bộ tiến trình thu thập và xử lý đã hoàn tất!")
}
