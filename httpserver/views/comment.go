package views

type GetComments struct {
	Id        int    `json:"id"`
	Message   string `json:"message"`
	PhotoId   int    `json:"photo_id"`
	UserId    int    `json:"user_id"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
