/*
Thử thách 5: Nút "Hủy Khẩn Cấp" (Worker Pool kết hợp Context)
Ngữ cảnh: Bạn có một hàng đợi công việc (Job Queue) chứa các mã đơn hàng cần xuất file. Các workers đang lấy việc ra làm. Đột nhiên có tín hiệu hủy bỏ chiến dịch từ hệ thống. Toàn bộ workers phải dừng ngay lập tức.

Yêu cầu kỹ thuật:

Thiết lập Worker Function:
Viết hàm worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup)

Trong hàm này, dùng vòng lặp vô hạn for { ... }.

Bên trong vòng lặp dùng lệnh select.

case <-ctx.Done(): In ra màn hình "Worker [id] nhận lệnh hủy. Đang dọn dẹp và thoát...", gọi wg.Done() rồi dùng return để kết thúc hoàn toàn hàm.

case job, ok := <-jobs: Nếu kênh chưa đóng, in ra "Worker [id] đang xử lý đơn hàng số [job]" và dùng time.Sleep(500 * time.Millisecond) để giả lập thời gian xử lý.

Thiết lập ở hàm main:

Tạo một jobs := make(chan int, 20) (kênh có bộ đệm).

Tạo một context có thể hủy: ctx, cancel := context.WithCancel(context.Background())

Tạo sync.WaitGroup và khởi chạy 3 workers.

Đẩy 20 ID đơn hàng (từ 1 đến 20) vào kênh jobs.

Kịch bản Hủy khẩn cấp:

Sau khi đẩy hết việc vào kênh, cho hàm main nghỉ 2 giây bằng time.Sleep(2 * time.Second). (Lúc này các workers sẽ kịp xử lý được vài đơn).

Bất ngờ gọi hàm cancel(). Lệnh này sẽ phát tín hiệu hủy qua ctx.Done() tới tất cả các workers cùng lúc.

Dùng wg.Wait() để đảm bảo tất cả workers đã thực sự dọn dẹp xong và báo cáo trước khi thoát chương trình.

Mục tiêu: Khi chạy code, bạn sẽ thấy các worker đang hì hục làm việc. Đúng 2 giây sau, lệnh cancel() được phát ra, tất cả đồng loạt quăng đồ nghề, in ra dòng chữ dọn dẹp và chương trình kết thúc ngay lập tức, bỏ mặc các đơn hàng còn tồn đọng trong kênh jobs.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d nhận lệnh hủy. Đang dọn dẹp và thoát...\n", id)
			wg.Done()
			return
		case job := <-jobs:
			fmt.Printf("Worker %d đang xử lý đơn hàng số %d\n", id, job)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	jobs := make(chan int, 20)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	var numberWorkers = 3
	wg.Add(numberWorkers)

	for i := 1; i <= numberWorkers; i++ {
		go worker(ctx, i, jobs, &wg)
	}
	for j := 1; j <= 20; j++ {
		jobs <- j
	}
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()

}
