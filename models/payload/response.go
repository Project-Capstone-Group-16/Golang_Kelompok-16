package payload

import "time"

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

type CreateFavoriteResponse struct {
	WarehouseID uint `json:"warehouse_id" form:"warehouse_id"`
	Warehouse   GetAllWarehouseResponse
}

// Admin Response
type GetAllWarehouseResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Status   string `json:"status"`
	Favorite uint   `json:"favorite"`
	ImageURL string `json:"image_url"`
}

type CreateAdminResponse struct {
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginAdminResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type CreateWarehouseResponse struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Status   string `json:"status"`
	ImageURL string `json:"image_url"`
}

type UpdateWarehouseResponse struct {
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Status   string `json:"status" form:"status"`
	ImageURL string `json:"image_url" form:"image_url"`
}

type ManageStaffResponse struct {
	FullName    string    `json:"full_name"`
	BirthDate   time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
}

type GetAllStaffsResponse struct {
	ID          uint      `json:"id"`
	WarehouseID uint      `json:warehouse_id"`
	FullName    string    `json:"full_name"`
	BirthDate   time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
}
