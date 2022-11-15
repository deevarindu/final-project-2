package params

type SocialMediaAddRequest struct {
	Name           string `validate:"required"`
	SocialMediaUrl string `validate:"required"`
}

type SocialMediaUpdateRequest struct {
	Name           string
	SocialMediaUrl string
}
