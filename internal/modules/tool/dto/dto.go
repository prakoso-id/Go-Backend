package dto

type CreateToolRequest struct {
	Name        string `json:"name" binding:"required"`
	Icon        string `json:"icon"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description"`
}

type CreateToolResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type UpdateToolRequest struct {
	Name        string `json:"name" binding:"required"`
	Icon        string `json:"icon"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description"`
}

type UpdateToolResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}

type ToolResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user,omitempty"`
}