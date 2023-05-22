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
