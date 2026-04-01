/*
Bài tập 3: Struct, Pointer và Method (Làm việc với Entity)
Yêu cầu: Go không có class, mà dùng struct để biểu diễn dữ liệu.

Tạo một struct tên là Product với các trường: ID (string), Name (string), Price (float64),
 và Quantity (int).

Viết một method tên là ApplyDiscount nhận vào phần trăm giảm giá (ví dụ: 10%)
và cập nhật lại giá (Price) của sản phẩm đó.

Trong hàm main(), khởi tạo một Product, in thông tin ra,
gọi hàm ApplyDiscount, và in lại thông tin để xem giá đã thực sự thay đổi chưa.

Gợi ý: Để thay đổi được giá trị của struct bên trong một method, bạn sẽ cần dùng đến Con trỏ (Pointer - *Product).
*/

package main

import "fmt"

type Product struct {
	ID       string
	name     string
	Price    float64
	Quantity int
}

func ApplyDiscount(entity *Product, percent float64) {
	entity.Price = entity.Price - entity.Price*percent/100

}

func main() {
	var p = Product{
		ID:       "1",
		name:     "Hàng hóa",
		Price:    300,
		Quantity: 20,
	}
	ApplyDiscount(&p, 30)
	fmt.Println(p)
}
