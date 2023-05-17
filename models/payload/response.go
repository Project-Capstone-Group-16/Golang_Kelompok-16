package payload

type CreateUserResponse struct {
	Email   string    `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}