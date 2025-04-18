package service

import (
	"errors"
	"go-backend/internal/modules/socialmedia/domain/entity"
	"go-backend/internal/modules/socialmedia/domain/repository"
	"go-backend/internal/modules/socialmedia/dto"
)

type SocialMediaService interface {
	Create(socialMedia *entity.SocialMedia) (*dto.CreateSocialMediaResponse, error)
	GetByID(id uint) (*dto.SocialMediaResponse, error)
	GetAll() ([]dto.SocialMediaResponse, error)
	Update(id uint, userID uint, req *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, error)
	Delete(id, userID uint) error
	GetByUserID(userID uint) ([]dto.SocialMediaResponse, error)
	GetByProfileID(profileID uint) ([]dto.SocialMediaResponse, error)
}

type socialMediaService struct {
	repo repository.SocialMediaRepository
}

func NewSocialMediaService(repo repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{repo: repo}
}

func (s *socialMediaService) Create(socialMedia *entity.SocialMedia) (*dto.CreateSocialMediaResponse, error) {
	if err := s.repo.Create(socialMedia); err != nil {
		return nil, err
	}

	return &dto.CreateSocialMediaResponse{
		ID:        socialMedia.ID,
		Platform:  socialMedia.Platform,
		Url:       socialMedia.Url,
		ProfileID: socialMedia.ProfileID,
		UserID:    socialMedia.UserID,
	}, nil
}

func (s *socialMediaService) GetByID(id uint) (*dto.SocialMediaResponse, error) {
	socialMedia, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.SocialMediaResponse{
		ID:        socialMedia.ID,
		Platform:  socialMedia.Platform,
		Url:       socialMedia.Url,
		ProfileID: socialMedia.ProfileID,
		UserID:    socialMedia.UserID,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    socialMedia.User.ID,
			Name:  socialMedia.User.Name,
			Email: socialMedia.User.Email,
		},
	}, nil
}

func (s *socialMediaService) GetAll() ([]dto.SocialMediaResponse, error) {
	socialMedias, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.SocialMediaResponse, len(socialMedias))
	for i, sm := range socialMedias {
		response[i] = dto.SocialMediaResponse{
			ID:        sm.ID,
			Platform:  sm.Platform,
			Url:       sm.Url,
			ProfileID: sm.ProfileID,
			UserID:    sm.UserID,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    sm.User.ID,
				Name:  sm.User.Name,
				Email: sm.User.Email,
			},
		}
	}

	return response, nil
}

func (s *socialMediaService) Update(id uint, userID uint, req *dto.UpdateSocialMediaRequest) (*dto.UpdateSocialMediaResponse, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if existing.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own social media")
	}

	existing.Platform = req.Platform
	existing.Url = req.Url

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return &dto.UpdateSocialMediaResponse{
		ID:        existing.ID,
		Platform:  existing.Platform,
		Url:       existing.Url,
		ProfileID: existing.ProfileID,
		UserID:    existing.UserID,
	}, nil
}

func (s *socialMediaService) Delete(id, userID uint) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only delete your own social media")
	}

	return s.repo.Delete(id)
}

func (s *socialMediaService) GetByUserID(userID uint) ([]dto.SocialMediaResponse, error) {
	socialMedias, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.SocialMediaResponse, len(socialMedias))
	for i, sm := range socialMedias {
		response[i] = dto.SocialMediaResponse{
			ID:        sm.ID,
			Platform:  sm.Platform,
			Url:       sm.Url,
			ProfileID: sm.ProfileID,
			UserID:    sm.UserID,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    sm.User.ID,
				Name:  sm.User.Name,
				Email: sm.User.Email,
			},
		}
	}

	return response, nil
}

func (s *socialMediaService) GetByProfileID(profileID uint) ([]dto.SocialMediaResponse, error) {
	socialMedias, err := s.repo.GetByProfileID(profileID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.SocialMediaResponse, len(socialMedias))
	for i, sm := range socialMedias {
		response[i] = dto.SocialMediaResponse{
			ID:        sm.ID,
			Platform:  sm.Platform,
			Url:       sm.Url,
			ProfileID: sm.ProfileID,
			UserID:    sm.UserID,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    sm.User.ID,
				Name:  sm.User.Name,
				Email: sm.User.Email,
			},
		}
	}

	return response, nil
}