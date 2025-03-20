package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// BaseResponse เป็นโครงสร้างมาตรฐานสำหรับการตอบกลับ API
type BaseResponse struct {
	Success   bool        `json:"success"`          // สถานะความสำเร็จ
	Code      int         `json:"code"`             // รหัส HTTP status
	Message   string      `json:"message"`          // ข้อความอธิบาย
	Data      interface{} `json:"data,omitempty"`   // ข้อมูลที่ส่งกลับ (ถ้ามี)
	Errors    interface{} `json:"errors,omitempty"` // รายละเอียดข้อผิดพลาด (ถ้ามี)
	Timestamp time.Time   `json:"timestamp"`        // เวลาที่ตอบกลับ
}

// MetaResponse สำหรับข้อมูล metadata เช่น pagination
type MetaResponse struct {
	Page        int `json:"page"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`
	CurrentPage int `json:"current_page"`
}

// PaginatedResponse สำหรับข้อมูลที่มีการแบ่งหน้า
type PaginatedResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Meta      MetaResponse `json:"meta"`
	Timestamp time.Time   `json:"timestamp"`
}

// NewSuccessResponse สร้าง BaseResponse สำหรับกรณีที่สำเร็จ
func NewSuccessResponse(data interface{}, message string) BaseResponse {
	if message == "" {
		message = "Success"
	}
	
	return BaseResponse{
		Success:   true,
		Code:      fiber.StatusOK,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewErrorResponse สร้าง BaseResponse สำหรับกรณีที่มีข้อผิดพลาด
func NewErrorResponse(code int, message string, errors interface{}) BaseResponse {
	if message == "" {
		message = "Error"
	}
	
	return BaseResponse{
		Success:   false,
		Code:      code,
		Message:   message,
		Errors:    errors,
		Timestamp: time.Now(),
	}
}

// NewCreatedResponse สร้าง BaseResponse สำหรับกรณีที่สร้างข้อมูลสำเร็จ
func NewCreatedResponse(data interface{}, message string) BaseResponse {
	if message == "" {
		message = "Created successfully"
	}
	
	return BaseResponse{
		Success:   true,
		Code:      fiber.StatusCreated,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse สร้าง PaginatedResponse
func NewPaginatedResponse(data interface{}, page, perPage, totalItems int, message string) PaginatedResponse {
	if message == "" {
		message = "Success"
	}
	
	totalPages := totalItems / perPage
	if totalItems%perPage != 0 {
		totalPages++
	}
	
	return PaginatedResponse{
		Success:   true,
		Code:      fiber.StatusOK,
		Message:   message,
		Data:      data,
		Meta: MetaResponse{
			Page:        page,
			PerPage:     perPage,
			TotalPages:  totalPages,
			TotalItems:  totalItems,
			CurrentPage: page,
		},
		Timestamp: time.Now(),
	}
}

// SendResponse ส่ง BaseResponse กลับไปยัง client
func SendResponse(c *fiber.Ctx, response BaseResponse) error {
	return c.Status(response.Code).JSON(response)
}

// SendPaginatedResponse ส่ง PaginatedResponse กลับไปยัง client
func SendPaginatedResponse(c *fiber.Ctx, response PaginatedResponse) error {
	return c.Status(response.Code).JSON(response)
}

// SendData ส่งข้อมูลที่สำเร็จกลับไปยัง client
func SendData(c *fiber.Ctx, data interface{}, message string) error {
	response := NewSuccessResponse(data, message)
	return c.Status(response.Code).JSON(response)
}

// SendError ส่งข้อผิดพลาดกลับไปยัง client
func SendError(c *fiber.Ctx, code int, message string, errors interface{}) error {
	response := NewErrorResponse(code, message, errors)
	return c.Status(code).JSON(response)
}

// SendCreated ส่งข้อมูลที่สร้างสำเร็จกลับไปยัง client
func SendCreated(c *fiber.Ctx, data interface{}, message string) error {
	response := NewCreatedResponse(data, message)
	return c.Status(response.Code).JSON(response)
}

// SendPaginated ส่งข้อมูลที่มีการแบ่งหน้ากลับไปยัง client
func SendPaginated(c *fiber.Ctx, data interface{}, page, perPage, totalItems int, message string) error {
	response := NewPaginatedResponse(data, page, perPage, totalItems, message)
	return c.Status(response.Code).JSON(response)
}