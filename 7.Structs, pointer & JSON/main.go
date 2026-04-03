/*
Thử thách 3: Structs, Pointers & JSON
Hệ thống nào cũng cần giao tiếp qua JSON.
Yêu cầu:

Tạo một struct ReportConfig gồm các trường: ReportName (string), IsActive (bool),
và TotalRecords (int). Lưu ý: Dùng struct tags (ví dụ `json:"report_name"`)
để format tên field khi chuyển sang JSON.

Viết một method Deactivate() cho ReportConfig để đổi IsActive thành false.
Lưu ý: Phải dùng Pointer Receiver để thay đổi được giá trị gốc.

Khởi tạo một ReportConfig, gọi Deactivate(), sau đó dùng package encoding/json
để Marshal struct đó thành chuỗi JSON và in ra màn hình.
*/
package main

import (
	"encoding/json"
	"fmt"
)

type ReportConfig struct {
	ReportName   string `json:"report_name"`
	IsActive     bool   `json:"is_active"`
	TotalRecords int    `json:"total_records"`
}

func (rc *ReportConfig) Deactivate() {
	rc.IsActive = false
}
func main() {
	ReportConfig := ReportConfig{
		ReportName:   "Báo cáo doanh thu",
		IsActive:     true,
		TotalRecords: 100,
	}
	ReportConfig.Deactivate()

	jsonData, err := json.MarshalIndent(ReportConfig, "", "  ")
	if err != nil {
		fmt.Println("Lỗi chuyển đổi JSON:", err)
		return
	}
	fmt.Println("\nTrạng thái LÚC SAU (Định dạng JSON):")
	fmt.Println(string(jsonData))

}
