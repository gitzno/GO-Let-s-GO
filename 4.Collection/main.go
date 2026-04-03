/*
Bài tập 4: Collection cốt lõi (Slice & Map)
Yêu cầu: Bạn có một danh sách các trạng thái đơn hàng:
statuses := []string{"pending", "shipped", "delivered", "pending", "canceled",
"shipped", "shipped"}

Hãy dùng một map[string]int để đếm số lượng của từng trạng thái
(Ví dụ: có bao nhiêu đơn "shipped", bao nhiêu đơn "pending").

In kết quả đếm ra màn hình.

Kiến thức cần dùng: Khởi tạo Slice, khởi tạo Map
(make), vòng lặp for ... range để duyệt qua collection.
*/

package main

import "fmt"

func main() {
	var statuses = []string{"pending", "shipped", "delivered", "pending", "canceled",
		"shipped", "shipped"}

	var mapFirst = make(map[string]int)

	for _, i := range statuses {
		mapFirst[i]++
	}
	for key, value := range mapFirst {
		fmt.Printf("%s count %d\n", key, value)
	}
}
