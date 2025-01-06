package dto

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreatePostResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type GetPostResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}