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
	ID        uint `json:"id"`
	Warehouse []models.Warehouse
}

// Admin Response
type GetAllWarehouseResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Capacity    uint   `json:"capacity"`
	Favorite    int    `json:"favorite"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
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
	ID          uint       `json:"id"`
	FullName    string     `json:"full_name"`
	Occupation  string     `json:"occupation" form:"occupation"`
	Gender      string     `json:"gender"`
	BirthDate   *time.Time `json:"birth_date"`
	PhoneNumber string     `json:"phone_number"`
	Address     string     `json:"address"`
	ImageURL    string     `json:"image_url"`
}

// Transaction Response

// type CreateTransactionResponse struct {
// 	UserID         uint `json:"user_id" form:"user_id"`
// 	User           User
// 	LockerID       uint `json:"locker_id" form:"locker_id"`
// 	Locker         Locker
// 	ItemCategoryID uint `json:"item_category_id" form:"item_category_id"`
// 	ItemCategory   ItemCategory
// 	Amount         uint       `json:"amount" form:"amount"`
// 	StartDate      *time.Time `json:"start_date" form:"start_date"`
// 	EndDate        *time.Time `json:"end_date" form:"end_date"`
// 	PaymentStatus  string     `json:"payment_status" form:"payment_status" gorm:"type:enum('Paid','Unpaid')"`
// 	PaymentUrl     string     `json:"payment_url" form:"payment_url"`
// }
