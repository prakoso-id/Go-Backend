package service

import (
	"fmt"
	experienceService "go-backend/internal/modules/experience/domain/service"
	postDTO "go-backend/internal/modules/post/dto"
	postService "go-backend/internal/modules/post/domain/service"
	profileService "go-backend/internal/modules/profile/domain/service"
	projectDTO "go-backend/internal/modules/project/dto"
	projectService "go-backend/internal/modules/project/domain/service"
	socialMediaDTO "go-backend/internal/modules/socialmedia/dto"
	socialMediaService "go-backend/internal/modules/socialmedia/domain/service"
	toolDTO "go-backend/internal/modules/tool/dto"
	toolService "go-backend/internal/modules/tool/domain/service"
	"go-backend/internal/modules/portfolio/dto"
)

type PortfolioService interface {
	GetUserPortfolio(userID uint) (*dto.PortfolioResponse, error)
	GetAllPortfolios() ([]*dto.PortfolioSummaryResponse, error)
}

type portfolioService struct {
	profileService    profileService.ProfileService
	postService       postService.PostService
	projectService    projectService.ProjectService
	socialMediaService socialMediaService.SocialMediaService
	toolService       toolService.ToolService
	experienceService experienceService.ExperienceService
}

func NewPortfolioService(
	profileSvc profileService.ProfileService,
	postSvc postService.PostService,
	projectSvc projectService.ProjectService,
	socialMediaSvc socialMediaService.SocialMediaService,
	toolSvc toolService.ToolService,
	experienceSvc experienceService.ExperienceService,
) PortfolioService {
	return &portfolioService{
		profileService:    profileSvc,
		postService:       postSvc,
		projectService:    projectSvc,
		socialMediaService: socialMediaSvc,
		toolService:       toolSvc,
		experienceService: experienceSvc,
	}
}

func (s *portfolioService) GetUserPortfolio(userID uint) (*dto.PortfolioResponse, error) {
	// Get user profile
	profiles, err := s.profileService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Check if profile exists
	if len(profiles) == 0 {
		return nil, fmt.Errorf("profile not found for user ID %d", userID)
	}

	// Use the first profile
	profile := &profiles[0]

	// Get user posts
	postsList, err := s.postService.ListByUserID(userID, 1, 100) // Using page 1 with 100 items per page
	if err != nil {
		return nil, err
	}
	
	// Convert slice of values to slice of pointers
	posts := make([]*postDTO.GetPostResponse, len(postsList))
	for i := range postsList {
		posts[i] = &postsList[i]
	}

	// Get user projects
	projectsList, err := s.projectService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert slice of project values to slice of pointers
	projects := make([]*projectDTO.ProjectResponse, len(projectsList))
	for i := range projectsList {
		projects[i] = &projectsList[i]
	}

	// Get user social media
	socialMediaList, err := s.socialMediaService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert slice of social media values to slice of pointers
	socialMedia := make([]*socialMediaDTO.SocialMediaResponse, len(socialMediaList))
	for i := range socialMediaList {
		socialMedia[i] = &socialMediaList[i]
	}

	// Get user tools
	toolsList, err := s.toolService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert slice of tool values to slice of pointers
	tools := make([]*toolDTO.ToolResponse, len(toolsList))
	for i := range toolsList {
		tools[i] = &toolsList[i]
	}

	// Get user experiences
	experiences, err := s.experienceService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.PortfolioResponse{
		Profile:     profile,
		Posts:       posts,
		Projects:    projects,
		SocialMedia: socialMedia,
		Tools:       tools,
		Experiences: experiences,
	}, nil
}

func (s *portfolioService) GetAllPortfolios() ([]*dto.PortfolioSummaryResponse, error) {
	// Get all profiles
	profiles, err := s.profileService.GetAll()
	if err != nil {
		return nil, err
	}

	var portfolioSummaries []*dto.PortfolioSummaryResponse
	for _, profile := range profiles {
		portfolioSummaries = append(portfolioSummaries, &dto.PortfolioSummaryResponse{
			UserID:      profile.UserID,
			Name:        profile.Name,
			ProfileImage: profile.ProfileImage,
			Bio:         profile.Bio,
		})
	}

	return portfolioSummaries, nil
}
