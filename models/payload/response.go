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

type GetAllWarehouseResponse struct {
	Name string `json:"name`
	Location string `json:"location"`
	Status string `json:"status"`
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
}

type UpdateWarehouseResponse struct {
	Name     string `json:"name" form:"name"`
	Location string `json:"location" form:"location"`
	Status   string `json:"status" form:"status"`
}
