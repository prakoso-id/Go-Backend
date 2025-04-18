package service

import (
	"errors"
	"go-backend/internal/modules/profile/domain/entity"
	"go-backend/internal/modules/profile/domain/repository"
	"go-backend/internal/modules/profile/dto"
	socialMediaDto "go-backend/internal/modules/socialmedia/dto"
)

type ProfileService interface {
	Create(profile *entity.Profile) (*dto.CreateProfileResponse, error)
	GetByID(id uint) (*dto.ProfileResponse, error)
	GetAll() ([]dto.ProfileResponse, error)
	Update(id uint, userID uint, req *dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error)
	Delete(id, userID uint) error
	GetByUserID(userID uint) ([]dto.ProfileResponse, error)
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) Create(profile *entity.Profile) (*dto.CreateProfileResponse, error) {
	if err := s.repo.Create(profile); err != nil {
		return nil, err
	}

	return &dto.CreateProfileResponse{
		ID:           profile.ID,
		Name:         profile.Name,
		Bio:          profile.Bio,
		ProfileImage: profile.ProfileImage,
		Email:        profile.Email,
		Phone:        profile.Phone,
		Location:     profile.Location,
		UserID:       profile.UserID,
	}, nil
}

func (s *profileService) GetByID(id uint) (*dto.ProfileResponse, error) {
	profile, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	socialMediaResponses := make([]socialMediaDto.SocialMediaResponse, len(profile.SocialMedia))
	for i, sm := range profile.SocialMedia {
		socialMediaResponses[i] = socialMediaDto.SocialMediaResponse{
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

	return &dto.ProfileResponse{
		ID:           profile.ID,
		Name:         profile.Name,
		Bio:          profile.Bio,
		ProfileImage: profile.ProfileImage,
		Email:        profile.Email,
		Phone:        profile.Phone,
		Location:     profile.Location,
		UserID:       profile.UserID,
		SocialMedia:  socialMediaResponses,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    profile.User.ID,
			Name:  profile.User.Name,
			Email: profile.User.Email,
		},
	}, nil
}

func (s *profileService) GetAll() ([]dto.ProfileResponse, error) {
	profiles, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ProfileResponse, len(profiles))
	for i, profile := range profiles {
		socialMediaResponses := make([]socialMediaDto.SocialMediaResponse, len(profile.SocialMedia))
		for j, sm := range profile.SocialMedia {
			socialMediaResponses[j] = socialMediaDto.SocialMediaResponse{
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

		response[i] = dto.ProfileResponse{
			ID:           profile.ID,
			Name:         profile.Name,
			Bio:          profile.Bio,
			ProfileImage: profile.ProfileImage,
			Email:        profile.Email,
			Phone:        profile.Phone,
			Location:     profile.Location,
			UserID:       profile.UserID,
			SocialMedia:  socialMediaResponses,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    profile.User.ID,
				Name:  profile.User.Name,
				Email: profile.User.Email,
			},
		}
	}

	return response, nil
}

func (s *profileService) Update(id uint, userID uint, req *dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if existing.UserID != userID {
		return nil, errors.New("unauthorized: you can only update your own profiles")
	}

	existing.Name = req.Name
	existing.Bio = req.Bio
	existing.ProfileImage = req.ProfileImage
	existing.Email = req.Email
	existing.Phone = req.Phone
	existing.Location = req.Location

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return &dto.UpdateProfileResponse{
		ID:           existing.ID,
		Name:         existing.Name,
		Bio:          existing.Bio,
		ProfileImage: existing.ProfileImage,
		Email:        existing.Email,
		Phone:        existing.Phone,
		Location:     existing.Location,
		UserID:       existing.UserID,
	}, nil
}

func (s *profileService) Delete(id, userID uint) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if existing.UserID != userID {
		return errors.New("unauthorized: you can only delete your own profiles")
	}

	return s.repo.Delete(id)
}

func (s *profileService) GetByUserID(userID uint) ([]dto.ProfileResponse, error) {
	profiles, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ProfileResponse, len(profiles))
	for i, profile := range profiles {
		socialMediaResponses := make([]socialMediaDto.SocialMediaResponse, len(profile.SocialMedia))
		for j, sm := range profile.SocialMedia {
			socialMediaResponses[j] = socialMediaDto.SocialMediaResponse{
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

		response[i] = dto.ProfileResponse{
			ID:           profile.ID,
			Name:         profile.Name,
			Bio:          profile.Bio,
			ProfileImage: profile.ProfileImage,
			Email:        profile.Email,
			Phone:        profile.Phone,
			Location:     profile.Location,
			UserID:       profile.UserID,
			SocialMedia:  socialMediaResponses,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    profile.User.ID,
				Name:  profile.User.Name,
				Email: profile.User.Email,
			},
		}
	}

	return response, nil
}