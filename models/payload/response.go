package payload

import (
	"Capstone/models"
	"time"
)

// User Response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type CreateUserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type GenerateOTPResponse struct {
	Email string `json:"email"`
}

type CreateFavoriteResponse struct {
	WarehouseID uint `json:"warehouse_id" form:"warehouse_id"`
	Warehouse   GetAllWarehouseResponse
}

type UpdateProfileUserResponse struct {
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	BirthDate   *time.Time `json:"birth_date"`
	Gender      string     `json:"gender"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	ImageURL    string     `json:"image_url"`
}

type FavoriteListUserResponse struct {
	ID        uint `json:"ID"`
	Warehouse []models.Warehouse
}

type GetAllUserResponse struct {
	ID                   uint       `json:"ID"`
	Email                string     `json:"email"`
	Fullname             string     `json:"fullname"`
	BirthDate            *time.Time `json:"birth_date"`
	PhoneNumber          string     `json:"phone_number"`
	Address              string     `json:"address"`
	Gender               string     `json:"gender"`
	TransactionHistroies int        `json:"transaction_histroies"`
	ImageURL             string     `json:"image_url"`
}

type GetExploreResponse struct {
	ActiveOrder      uint `json:"active_order"`
	CountTransaction uint `json:"count_transaction"`
}

// Admin Response
type GetAllWarehouseResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Capacity    uint   `json:"capacity"`
	Favorite    int    `json:"favorite"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Locker      []models.Locker
}

type CreateAdminResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginAdminResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type CreateWarehouseResponse struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Capacity    uint   `json:"capacity"`
	Status      string `json:"status"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type UpdateWarehouseResponse struct {
	Name        string `json:"name" form:"name"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Capacity    uint   `json:"capacity"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url" form:"image_url"`
}

type ManageStaffResponse struct {
	FullName    string     `json:"full_name"`
	Occupation  string     `json:"occupation" form:"occupation"`
	Gender      string     `json:"gender" `
	BirthDate   *time.Time `json:"birth_date"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	ImageURL    string     `json:"image_url"`
}

type GetAllStaffsResponse struct {
	ID          uint       `json:"ID"`
	FullName    string     `json:"full_name"`
	Occupation  string     `json:"occupation" form:"occupation"`
	Gender      string     `json:"gender"`
	BirthDate   *time.Time `json:"birth_date"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	ImageURL    string     `json:"image_url"`
}

// Transaction Response

type MidtransStatusResponse struct {
	OrderID           string `json:"order_id"`
	TransactionID     string `json:"transaction_id"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
}

//Dashboard admin Response

type DashboardAdminResponse struct {
	Todey            time.Time `json:"todey"`
	TotalLockers     uint      `json:"total_lockers"`
	TotalUsedLockers uint      `json:"total_used_lockers"`
	TotalUsers       uint      `json:"total_users"`
	TotalIncome      uint      `json:"total_income"`
}
