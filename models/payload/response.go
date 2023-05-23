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