package dto

import (
	"encoding/json"
	"go-backend/internal/modules/experience/domain/entity"
	"time"
)

// CreateExperienceRequest represents the request for creating a new experience
type CreateExperienceRequest struct {
	Title       string    `json:"title" binding:"required"`
	Company     string    `json:"company" binding:"required"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date"`
	Description string    `json:"description"`
	TechStack   []string  `json:"tech_stack"`
}

// UpdateExperienceRequest represents the request for updating an experience
type UpdateExperienceRequest struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Description string    `json:"description"`
	TechStack   []string  `json:"tech_stack"`
}

// ExperienceResponse represents the response for an experience
type ExperienceResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Description string    `json:"description"`
	TechStack   []string  `json:"tech_stack"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToEntity converts a CreateExperienceRequest to an Experience entity
func (r *CreateExperienceRequest) ToEntity(userID uint) (*entity.Experience, error) {
	techStackJSON, err := json.Marshal(r.TechStack)
	if err != nil {
		return nil, err
	}
	
	return &entity.Experience{
		Title:       r.Title,
		Company:     r.Company,
		Location:    r.Location,
		StartDate:   r.StartDate,
		EndDate:     r.EndDate,
		Description: r.Description,
		TechStack:   string(techStackJSON),
		UserID:      userID,
	}, nil
}

// ToResponse converts an Experience entity to an ExperienceResponse
func ToResponse(experience *entity.Experience) (*ExperienceResponse, error) {
	var techStack []string
	if experience.TechStack != "" {
		if err := json.Unmarshal([]byte(experience.TechStack), &techStack); err != nil {
			return nil, err
		}
	}

	return &ExperienceResponse{
		ID:          experience.ID,
		Title:       experience.Title,
		Company:     experience.Company,
		Location:    experience.Location,
		StartDate:   experience.StartDate,
		EndDate:     experience.EndDate,
		Description: experience.Description,
		TechStack:   techStack,
		UserID:      experience.UserID,
		CreatedAt:   experience.CreatedAt,
		UpdatedAt:   experience.UpdatedAt,
	}, nil
}

// ToResponseList converts a slice of Experience entities to a slice of ExperienceResponse
func ToResponseList(experiences []entity.Experience) ([]*ExperienceResponse, error) {
	var responseList []*ExperienceResponse
	for _, experience := range experiences {
		exp := experience // Create a copy to avoid issues with loop variable
		response, err := ToResponse(&exp)
		if err != nil {
			return nil, err
		}
		responseList = append(responseList, response)
	}
	return responseList, nil
}

// UpdateEntity updates an Experience entity with the values from UpdateExperienceRequest
func (r *UpdateExperienceRequest) UpdateEntity(experience *entity.Experience) error {
	if r.Title != "" {
		experience.Title = r.Title
	}
	if r.Company != "" {
		experience.Company = r.Company
	}
	if r.Location != "" {
		experience.Location = r.Location
	}
	if !r.StartDate.IsZero() {
		experience.StartDate = r.StartDate
	}
	if r.EndDate != nil {
		experience.EndDate = r.EndDate
	}
	if r.Description != "" {
		experience.Description = r.Description
	}
	if r.TechStack != nil {
		techStackJSON, err := json.Marshal(r.TechStack)
		if err != nil {
			return err
		}
		experience.TechStack = string(techStackJSON)
	}
	return nil
}
