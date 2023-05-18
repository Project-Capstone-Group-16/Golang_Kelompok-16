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
