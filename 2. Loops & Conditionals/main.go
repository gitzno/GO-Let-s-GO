/*
Bài tập 2: Cấu trúc điều khiển (Loops & Conditionals)
Yêu cầu: Go chỉ có duy nhất một vòng lặp là for. Hãy dùng nó để giải quyết bài toán sau:

Viết một hàm isPrime(n int) bool để kiểm tra xem một số nguyên n có phải là số nguyên tố hay không.

Trong hàm main(), dùng vòng lặp để in ra tất cả các số nguyên tố từ 1 đến 50.

Kiến thức cần dùng: Vòng lặp for, câu lệnh điều kiện if/else, cách khai báo và trả về giá trị từ function.
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 || n == 4 {
		return false
	}
	if n == 2 || n == 3 || n == 5 {
		return true
	}
	// fmt.Printf("%d = %d \n", n, int(math.Sqrt(float64(n))))
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	print("Các số nguyên tố là: ")
	for i := 0; i <= 50; i++ {
		if isPrime(i) {
			fmt.Printf(" %d ", i)
		}
	}
}
