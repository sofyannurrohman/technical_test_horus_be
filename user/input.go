package user

type RegisterUserInput struct{
	UserName string `json:"username" binding:"required"`
	Name string 	`json:"name" binding:"required"`
	Email string	`json:"email" binding:"required,email"`
	Password string`json:"password" binding:"required"`
}
type LoginInput struct{
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct{
	Email string `json:"email" binding:"required,email"`
}