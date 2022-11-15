package params

type UserCreateRequest struct {
	Username string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Age      int    `validate:"required"`
}

type UserLoginRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UserUpdateRequest struct {
	Username string
	Email    string
	Password string
	Age      int
}
