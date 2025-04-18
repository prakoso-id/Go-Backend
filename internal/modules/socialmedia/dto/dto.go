package dto

type CreateSocialMediaRequest struct {
	Platform  string `json:"platform" binding:"required"`
	Url       string `json:"url" binding:"required"`
	ProfileID uint   `json:"profile_id" binding:"required"`
}

type CreateSocialMediaResponse struct {
	ID        uint   `json:"id"`
	Platform  string `json:"platform"`
	Url       string `json:"url"`
	ProfileID uint   `json:"profile_id"`
	UserID    uint   `json:"user_id"`
}

type UpdateSocialMediaRequest struct {
	Platform string `json:"platform" binding:"required"`
	Url      string `json:"url" binding:"required"`
}

type UpdateSocialMediaResponse struct {
	ID        uint   `json:"id"`
	Platform  string `json:"platform"`
	Url       string `json:"url"`
	ProfileID uint   `json:"profile_id"`
	UserID    uint   `json:"user_id"`
}

type SocialMediaResponse struct {
	ID        uint   `json:"id"`
	Platform  string `json:"platform"`
	Url       string `json:"url"`
	ProfileID uint   `json:"profile_id"`
	UserID    uint   `json:"user_id"`
	User      struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user,omitempty"`
}