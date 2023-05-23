package payload

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

// Admin Response

type CreateAdminResponse struct {
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
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
