package dto

import (
	socialMediaDto "go-backend/internal/modules/socialmedia/dto"
)

type CreateProfileRequest struct {
	Name         string `json:"name" binding:"required"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
}

type CreateProfileResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
	UserID       uint   `json:"user_id"`
}

type UpdateProfileRequest struct {
	Name         string `json:"name" binding:"required"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
}

type UpdateProfileResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Bio          string `json:"bio"`
	ProfileImage string `json:"profile_image"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
	UserID       uint   `json:"user_id"`
}

type ProfileResponse struct {
	ID           uint                             `json:"id"`
	Name         string                           `json:"name"`
	Bio          string                           `json:"bio"`
	ProfileImage string                           `json:"profile_image"`
	Email        string                           `json:"email"`
	Phone        string                           `json:"phone"`
	Location     string                           `json:"location"`
	UserID       uint                             `json:"user_id"`
	SocialMedia  []socialMediaDto.SocialMediaResponse `json:"social_media,omitempty"`
	User         struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user,omitempty"`
}