package service

import (
	"errors"
	imageEntity "go-backend/internal/modules/images/domain/entity"
	postEntity "go-backend/internal/modules/post/domain/entity"
	"go-backend/internal/modules/post/domain/repository"
	"go-backend/internal/modules/post/dto"
)

type PostService interface {
	Create(userID uint, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error)
	GetByID(id uint) (*dto.GetPostResponse, error)
	Update(id, userID uint, req *dto.UpdatePostRequest) (*dto.UpdatePostResponse, error)
	Delete(id, userID uint) error
	List(page, pageSize int) ([]dto.GetPostResponse, error)
	ListByUserID(userID uint, page, pageSize int) ([]dto.GetPostResponse, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) Create(userID uint, req *dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	post := &postEntity.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	// Create images if provided
	if len(req.ImageURLs) > 0 {
		images := make([]imageEntity.Images, len(req.ImageURLs))
		for i, url := range req.ImageURLs {
			images[i] = imageEntity.Images{
				URL:    url,
				UserID: userID,
			}
		}
		post.Images = images
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(post.Images))
	for i, img := range post.Images {
		imageURLs[i] = img.URL
	}

	return &dto.CreatePostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		ImageURLs: imageURLs,
	}, nil
}

func (s *postService) GetByID(id uint) (*dto.GetPostResponse, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(post.Images))
	for i, img := range post.Images {
		imageURLs[i] = img.URL
	}

	resp := &dto.GetPostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		ImageURLs: imageURLs,
		User: struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
	}

	return resp, nil
}

func (s *postService) Update(id, userID uint, req *dto.UpdatePostRequest) (*dto.UpdatePostResponse, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if post.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	// Update images if provided
	if len(req.ImageURLs) > 0 {
		images := make([]imageEntity.Images, len(req.ImageURLs))
		for i, url := range req.ImageURLs {
			images[i] = imageEntity.Images{
				URL:    url,
				UserID: userID,
			}
		}
		post.Images = images
	}

	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	// Extract image URLs for response
	imageURLs := make([]string, len(post.Images))
	for i, img := range post.Images {
		imageURLs[i] = img.URL
	}

	return &dto.UpdatePostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserID:    post.UserID,
		ImageURLs: imageURLs,
	}, nil
}

func (s *postService) Delete(id, userID uint) error {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}

func (s *postService) List(page, pageSize int) ([]dto.GetPostResponse, error) {
	offset := (page - 1) * pageSize
	posts, err := s.repo.List(offset, pageSize)
	if err != nil {
		return nil, err
	}

	response := make([]dto.GetPostResponse, len(posts))
	for i, post := range posts {
		// Extract image URLs for response
		imageURLs := make([]string, len(post.Images))
		for j, img := range post.Images {
			imageURLs[j] = img.URL
		}

		response[i] = dto.GetPostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			ImageURLs: imageURLs,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
		}
	}

	return response, nil
}

func (s *postService) ListByUserID(userID uint, page, pageSize int) ([]dto.GetPostResponse, error) {
	offset := (page - 1) * pageSize
	posts, err := s.repo.ListByUserID(userID, offset, pageSize)
	if err != nil {
		return nil, err
	}

	response := make([]dto.GetPostResponse, len(posts))
	for i, post := range posts {
		// Extract image URLs for response
		imageURLs := make([]string, len(post.Images))
		for j, img := range post.Images {
			imageURLs[j] = img.URL
		}

		response[i] = dto.GetPostResponse{
			ID:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserID:    post.UserID,
			ImageURLs: imageURLs,
			User: struct {
				ID    uint   `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
		}
	}

	return response, nil
}