package dto

type CreateProjectRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	ImageURLs   []string `json:"image_urls,omitempty"`
}

type CreateProjectResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	UserID      uint     `json:"user_id"`
	ImageURLs   []string `json:"image_urls,omitempty"`
}

type UpdateProjectRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	ImageURLs   []string `json:"image_urls,omitempty"`
}

type UpdateProjectResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	UserID      uint     `json:"user_id"`
	ImageURLs   []string `json:"image_urls,omitempty"`
}

type ProjectResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	UserID      uint     `json:"user_id"`
	ImageURLs   []string `json:"image_urls,omitempty"`
	User        struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user,omitempty"`
}
