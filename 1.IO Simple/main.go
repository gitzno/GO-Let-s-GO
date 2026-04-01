/*
Bài tập 1: Nhập xuất cơ bản và Biến (I/O & Variables)

Yêu cầu: Viết một chương trình chạy trên Terminal (CLI).

Chương trình in ra câu hỏi yêu cầu người dùng nhập vào Tên và Năm sinh.

Đọc dữ liệu người dùng nhập vào.

Tính tuổi của người đó (dựa trên năm hiện tại) và in ra màn hình một câu chào chuẩn form. Ví dụ: "Chào John, năm nay bạn 25 tuổi!"

Kiến thức cần dùng: Package fmt (fmt.Println, fmt.Scan hoặc bufio), cách khai báo biến (var, :=).
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var name string
	var year int
	fmt.Println("Hello người dùng!")
	fmt.Print("Nhập tên của bạn:")
	fmt.Scan(&name)
	fmt.Print("Nhập năm của bạn:")
	fmt.Scan(&year)
	var now = time.Now().Local().Year()
	fmt.Printf("Chào %s tuổi của bạn là %d", name, now-year)

	// var age = strconv.ParseInt(now.Format("2002"), 10, 64) - year
	// fmt.Printf("Thời gian hiện tại %d", age)

}
