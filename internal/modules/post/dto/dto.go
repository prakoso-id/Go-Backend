package dto

type CreatePostRequest struct {
	Title    string   `json:"title" binding:"required"`
	Content  string   `json:"content" binding:"required"`
	ImageURLs []string `json:"image_urls"`
}

type CreatePostResponse struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	UserID    uint     `json:"user_id"`
	ImageURLs []string `json:"image_urls"`
}

type UpdatePostRequest struct {
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	ImageURLs []string `json:"image_urls"`
}

type UpdatePostResponse struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	UserID    uint     `json:"user_id"`
	ImageURLs []string `json:"image_urls"`
}

type GetPostResponse struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	UserID    uint     `json:"user_id"`
	ImageURLs []string `json:"image_urls"`
	User      struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}