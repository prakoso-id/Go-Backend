package service

import (
	"go-backend/internal/modules/experience/domain/repository"
	"go-backend/internal/modules/experience/dto"
)

type ExperienceService interface {
	Create(request *dto.CreateExperienceRequest, userID uint) (*dto.ExperienceResponse, error)
	GetAll() ([]*dto.ExperienceResponse, error)
	GetByID(id uint) (*dto.ExperienceResponse, error)
	GetByUserID(userID uint) ([]*dto.ExperienceResponse, error)
	Update(id uint, request *dto.UpdateExperienceRequest) (*dto.ExperienceResponse, error)
	Delete(id uint) error
}

type experienceService struct {
	repo repository.ExperienceRepository
}

func NewExperienceService(repo repository.ExperienceRepository) ExperienceService {
	return &experienceService{
		repo: repo,
	}
}

func (s *experienceService) Create(request *dto.CreateExperienceRequest, userID uint) (*dto.ExperienceResponse, error) {
	experience, err := request.ToEntity(userID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(experience); err != nil {
		return nil, err
	}

	return dto.ToResponse(experience)
}

func (s *experienceService) GetAll() ([]*dto.ExperienceResponse, error) {
	experiences, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return dto.ToResponseList(experiences)
}

func (s *experienceService) GetByID(id uint) (*dto.ExperienceResponse, error) {
	experience, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToResponse(experience)
}

func (s *experienceService) GetByUserID(userID uint) ([]*dto.ExperienceResponse, error) {
	experiences, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return dto.ToResponseList(experiences)
}



func (s *experienceService) Update(id uint, request *dto.UpdateExperienceRequest) (*dto.ExperienceResponse, error) {
	experience, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if err := request.UpdateEntity(experience); err != nil {
		return nil, err
	}

	if err := s.repo.Update(experience); err != nil {
		return nil, err
	}

	return dto.ToResponse(experience)
}

func (s *experienceService) Delete(id uint) error {
	return s.repo.Delete(id)
}
