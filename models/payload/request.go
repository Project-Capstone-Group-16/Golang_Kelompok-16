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
	Email string `json:"email" form:"email" validate:"required,email"`
	Otp   string `json:"otp" form:"otp" validate:"required,min=4"`
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
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"max=11,min=10,number"`
	Address     string `json:"address" form:"address"`
	ImageURL    string `json:"image_url" form:"image_url"`
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
	Name        string `json:"name" form:"name" validate:"required"`
	City        string `json:"city" form:"city" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	ImageURL    string `json:"image_url" form:"image_url" validate:"required"`
}

type UpdateWarehouseRequest struct {
	Name           string `json:"name" form:"name"`
	City           string `json:"city" form:"city"`
	Address        string `json:"address" form:"address"`
	Status         string `json:"status" form:"status"`
	Description    string `json:"description" form:"description" validate:"required"`
	WarehouseImage string `json:"warehouse_image" form:"warehouse_image"`
}

type CreateStaffRequest struct {
	FullName    string `json:"full_name" form:"full_name" validate:"required"`
	Occupation  string `json:"occupation" form:"occupation" validate:"required"`
	Gender      string `json:"gender" form:"gender" gorm:"type:enum('PRIA', 'WANITA')" validate:"required"`
	BirthDate   string `json:"birth_date" form:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required,max=12,min=11,number"`
	Address     string `json:"address" form:"address" validate:"required"`
	ImageURL    string `json:"image_url" form:"image_url" validate:"required"`
}

type UpdateStaffRequest struct {
	FullName    string `json:"full_name" form:"full_name"`
	Occupation  string `json:"occupation" form:"occupation"`
	Gender      string `json:"gender" form:"gender" gorm:"type:enum('PRIA', 'WANITA')" validate:"required"`
	BirthDate   string `json:"birth_date" form:"birth_date"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"omitempty,gt=0,max=12,min=11,number"`
	Address     string `json:"address" form:"address"`
	ImageURL    string `json:"image_url" form:"image_url"`
}

//  Transaction Request
type CreateTransactionRequest struct {
	UserID         uint   `json:"user_id" form:"user_id"`
	LockerTypeID   uint   `json:"locker_type_id" form:"locker_type_id" validate:"required"`
	ItemCategoryID uint   `json:"item_category_id" form:"item_category_id" validate:"required"`
	WarehouseID    uint   `json:"warehouse_id" form:"warehouse_id" validate:"required"`
	StartDate      string `json:"start_date" form:"start_date" validate:"required"`
	EndDate        string `json:"end_date" form:"end_date" validate:"required"`
	PaymentStatus  string `json:"payment_status" form:"payment_status" gorm:"type:enum('Paid','Unpaid')"`
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	TransactionTime   string `json:"transaction_time"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
}
