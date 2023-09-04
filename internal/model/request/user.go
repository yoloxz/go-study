package request

type Register struct {
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	PassWord string `json:"password" binding:"required"`
}
