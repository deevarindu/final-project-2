package params

type PhotoUploadRequest struct {
	Title    string `validate:"required"`
	Caption  string `validate:"required"`
	PhotoUrl string `validate:"required"`
	UserId   int    `validate:"required"`
}

type PhotoUpdateRequest struct {
	Title    string
	Caption  string
	PhotoUrl string
}
