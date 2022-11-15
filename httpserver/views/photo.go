package views

type GetPhotos struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
