package views

type GetUsers struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type GetUser struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
