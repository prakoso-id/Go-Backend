package service

import (
	"errors"
	imageEntity "go-backend/internal/modules/images/domain/entity"
	"go-backend/internal/modules/project/domain/entity"
	"go-backend/internal/modules/project/domain/repository"
	"go-backend/internal/modules/project/dto"
)

type ProjectService interface {
	Create(project *entity.Project) (*dto.CreateProjectResponse, error)
	GetByID(id uint) (*dto.ProjectResponse, error)
	GetAll() ([]dto.ProjectResponse, error)
	Update(id uint, userID uint, req *dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error)
	Delete(id, userID uint) error
	GetByUserID(userID uint) ([]dto.ProjectResponse, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) Create(project *entity.Project) (*dto.CreateProjectResponse, error) {
	// Create images if provided
	if len(project.Images) > 0 {
		for i := range project.Images {
			project.Images[i].UserID = project.UserID
		}
	}

	if err := s.repo.Create(project); err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(project.Images))
	for i, img := range project.Images {
		imageURLs[i] = img.URL
	}

	return &dto.CreateProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Url:         project.Url,
		UserID:      project.UserID,
		ImageURLs:   imageURLs,
	}, nil
}

func (s *projectService) GetByID(id uint) (*dto.ProjectResponse, error) {
	project, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(project.Images))
	for i, img := range project.Images {
		imageURLs[i] = img.URL
	}

	resp := &dto.ProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Url:         project.Url,
		UserID:      project.UserID,
		ImageURLs:   imageURLs,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    project.User.ID,
			Name:  project.User.Name,
			Email: project.User.Email,
		},
	}

	return resp, nil
}

func (s *projectService) GetAll() ([]dto.ProjectResponse, error) {
	projects, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ProjectResponse, len(projects))
	for i, project := range projects {
		// Extract image URLs for response
		imageURLs := make([]string, len(project.Images))
		for j, img := range project.Images {
			imageURLs[j] = img.URL
		}

		response[i] = dto.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Url:         project.Url,
			UserID:      project.UserID,
			ImageURLs:   imageURLs,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    project.User.ID,
				Name:  project.User.Name,
				Email: project.User.Email,
			},
		}
	}

	return response, nil
}

func (s *projectService) Update(id uint, userID uint, req *dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error) {
	project, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if project.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own projects")
	}

	project.Name = req.Name
	project.Description = req.Description
	project.Url = req.Url

	// Update images if provided
	if len(req.ImageURLs) > 0 {
		images := make([]imageEntity.Images, len(req.ImageURLs))
		for i, url := range req.ImageURLs {
			images[i] = imageEntity.Images{
				URL:       url,
				UserID:    userID,
			}
		}
		project.Images = images
	}

	if err := s.repo.Update(project); err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(project.Images))
	for i, img := range project.Images {
		imageURLs[i] = img.URL
	}

	return &dto.UpdateProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		Url:         project.Url,
		UserID:      project.UserID,
		ImageURLs:   imageURLs,
	}, nil
}

func (s *projectService) Delete(id, userID uint) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only delete your own projects")
	}

	return s.repo.Delete(id)
}

func (s *projectService) GetByUserID(userID uint) ([]dto.ProjectResponse, error) {
	projects, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ProjectResponse, len(projects))
	for i, project := range projects {
		// Extract image URLs for response
		imageURLs := make([]string, len(project.Images))
		for j, img := range project.Images {
			imageURLs[j] = img.URL
		}

		response[i] = dto.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			Url:         project.Url,
			UserID:      project.UserID,
			ImageURLs:   imageURLs,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    project.User.ID,
				Name:  project.User.Name,
				Email: project.User.Email,
			},
		}
	}

	return response, nil
}