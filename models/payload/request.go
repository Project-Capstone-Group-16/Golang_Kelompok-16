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

type VerifyngOtpRequest struct {
	Otp string `json:"otp" validate:"required,min=4"`
}

type CreateFavoriteRequest struct {
	UserID      uint `json:"user_id" form:"user_id"`
	WarehouseID uint `json:"warehouse_id" form:"warehouse_id" validate:"required"`
}

type UpdateProfileUser struct {
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	BirthDate   string `json:"birth_date" form:"birth_date"`
	Gender      string `json:"gender" form:"gender" gorm:"type:enum('PRIA', 'WANITA', '');default:''"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=11,min=10,number`
	Address     string `json:"address" form:"address"`
}

// admin request

type CreateAdminRequest struct {
	FirstName       string `json:"first_name" form:"first_name" validate:"required"`
	LastName        string `json:"last_name" form:"last_name" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" validate:"required,max=11,min=10,number"`
	Password        string `json:"password" form:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginAdminRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

type CreateWarehouseRequest struct {
	Name           string `json:"name" form:"name" validate:"required"`
	Location       string `json:"location" form:"location" validate:"required"`
	WarehouseImage string `json:"warehouse_image" form:"warehouse_image" validate:"required"`
}

type CreateStaffRequest struct {
	FullName     string `json:"full_name" form:"full_name" validate:"required"`
	WarehouseID  uint   `json:"warehouse_id" form:"warehouse_id" validate:"required"`
	Occupation   string `json:"occupation" form:"occupation" validate:"required"`
	BirthDate    string `json:"birth_date" form:"birth_date" validate:"required"`
	PhoneNumber  string `json:"phone_number" form:"phone_number" validate:"required,max=11,min=10,number"`
}

type UpdateStaffRequest struct {
	FullName     string `json:"full_name" form:"full_name"`
	WarehouseID  uint   `json:"warehouse_id" form:"warehouse_id"`
	Occupation   string `json:"occupation" form:"occupation"`
	BirthDate    string `json:"birth_date" form:"birth_date"`
	PhoneNumber  string `json:"phone_number" form:"phone_number" validate:"omitempty,gt=0,max=11,min=10,number"`
}
