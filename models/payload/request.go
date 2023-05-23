package payload

// User Request
type CreateUserRequest struct {
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

type UpdatePasswordRequest struct {
	Password        string `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required"`
}

type VerifyngOtp struct {
	Otp string `json:"otp" validate:"required,min=6"`
}

// admin request

type CreateAdminRequest struct {
	Fullname        string `json:"full_name" form:"full_name" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	PhoneNumber     int    `json:"phone_number" form:"phone_number" validate:"required,min=11"`
	Password        string `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginAdminRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type CreateWarehouseRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Location string `json:"location" form:"location" validate:"required"`
	WarehouseImage string `json:"warehouse_image" form:"warehouse_image" validate:"required"`
}

