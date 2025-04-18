package dto

import (
	experienceDTO "go-backend/internal/modules/experience/dto"
	postDTO "go-backend/internal/modules/post/dto"
	profileDTO "go-backend/internal/modules/profile/dto"
	projectDTO "go-backend/internal/modules/project/dto"
	socialMediaDTO "go-backend/internal/modules/socialmedia/dto"
	toolDTO "go-backend/internal/modules/tool/dto"
)

// PortfolioResponse represents a complete user portfolio with all associated data
type PortfolioResponse struct {
	Profile     *profileDTO.ProfileResponse              `json:"profile"`
	Posts       []*postDTO.GetPostResponse               `json:"posts"`
	Projects    []*projectDTO.ProjectResponse            `json:"projects"`
	SocialMedia []*socialMediaDTO.SocialMediaResponse    `json:"social_media"`
	Tools       []*toolDTO.ToolResponse                  `json:"tools"`
	Experiences []*experienceDTO.ExperienceResponse      `json:"experiences"`
}

// PortfolioSummaryResponse represents a summary of a user's portfolio for listing
type PortfolioSummaryResponse struct {
	UserID       uint   `json:"user_id"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	Bio          string `json:"bio"`
}
