package request

type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Rusult  string `json:"result"`
}

type Register struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	PassWord string `json:"password" binding:"required"`
}
