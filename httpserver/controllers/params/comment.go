package params

type CommentCreateRequest struct {
	Message string `validate:"required"`
	PhotoId int    `validate:"required"`
}

type CommentUpdateRequest struct {
	Message string
}
