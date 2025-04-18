package service

import (
	"errors"
	"go-backend/internal/modules/tool/domain/entity"
	"go-backend/internal/modules/tool/domain/repository"
	"go-backend/internal/modules/tool/dto"
)

type ToolService interface {
	Create(tool *entity.Tool) (*dto.CreateToolResponse, error)
	GetByID(id uint) (*dto.ToolResponse, error)
	GetAll() ([]dto.ToolResponse, error)
	Update(id uint, userID uint, req *dto.UpdateToolRequest) (*dto.UpdateToolResponse, error)
	Delete(id, userID uint) error
	GetByUserID(userID uint) ([]dto.ToolResponse, error)
}

type toolService struct {
	repo repository.ToolRepository
}

func NewToolService(repo repository.ToolRepository) ToolService {
	return &toolService{repo: repo}
}

func (s *toolService) Create(tool *entity.Tool) (*dto.CreateToolResponse, error) {
	if err := s.repo.Create(tool); err != nil {
		return nil, err
	}

	return &dto.CreateToolResponse{
		ID:          tool.ID,
		Name:        tool.Name,
		Icon:        tool.Icon,
		Category:    tool.Category,
		Description: tool.Description,
		UserID:      tool.UserID,
	}, nil
}

func (s *toolService) GetByID(id uint) (*dto.ToolResponse, error) {
	tool, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := &dto.ToolResponse{
		ID:          tool.ID,
		Name:        tool.Name,
		Icon:        tool.Icon,
		Category:    tool.Category,
		Description: tool.Description,
		UserID:      tool.UserID,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    tool.User.ID,
			Name:  tool.User.Name,
			Email: tool.User.Email,
		},
	}

	return resp, nil
}

func (s *toolService) GetAll() ([]dto.ToolResponse, error) {
	tools, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ToolResponse, len(tools))
	for i, tool := range tools {
		response[i] = dto.ToolResponse{
			ID:          tool.ID,
			Name:        tool.Name,
			Icon:        tool.Icon,
			Category:    tool.Category,
			Description: tool.Description,
			UserID:      tool.UserID,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    tool.User.ID,
				Name:  tool.User.Name,
				Email: tool.User.Email,
			},
		}
	}

	return response, nil
}

func (s *toolService) Update(id uint, userID uint, req *dto.UpdateToolRequest) (*dto.UpdateToolResponse, error) {
	tool, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if tool.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own tools")
	}

	tool.Name = req.Name
	tool.Icon = req.Icon
	tool.Category = req.Category
	tool.Description = req.Description

	if err := s.repo.Update(tool); err != nil {
		return nil, err
	}

	return &dto.UpdateToolResponse{
		ID:          tool.ID,
		Name:        tool.Name,
		Icon:        tool.Icon,
		Category:    tool.Category,
		Description: tool.Description,
		UserID:      tool.UserID,
	}, nil
}

func (s *toolService) Delete(id, userID uint) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only delete your own tools")
	}

	return s.repo.Delete(id)
}

func (s *toolService) GetByUserID(userID uint) ([]dto.ToolResponse, error) {
	tools, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ToolResponse, len(tools))
	for i, tool := range tools {
		response[i] = dto.ToolResponse{
			ID:          tool.ID,
			Name:        tool.Name,
			Icon:        tool.Icon,
			Category:    tool.Category,
			Description: tool.Description,
			UserID:      tool.UserID,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    tool.User.ID,
				Name:  tool.User.Name,
				Email: tool.User.Email,
			},
		}
	}

	return response, nil
}