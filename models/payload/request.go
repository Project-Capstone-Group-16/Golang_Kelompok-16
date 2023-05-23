package payload

// User Request
type CreateUserRequest struct {
	Email           string `json:"email" form:"email" validate:"required,email"`
	Password        string `json:"password" form:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" form:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type VerifyngOtp struct {
	Otp string `json:"otp" validate:"required,min=6"`
}

type DeleteWarehouseRequest struct{
	WarehouseID uint `json:"warehouse_id" form:"warehouse_id"`
}