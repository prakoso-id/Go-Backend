package service

import (
	"errors"
	"go-backend/internal/modules/post/domain/entity"
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
	post := &entity.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	return &dto.CreatePostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func (s *postService) GetByID(id uint) (*dto.GetPostResponse, error) {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	resp := &dto.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
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
		return nil, errors.New("unauthorized: you can only update your own posts")
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	return &dto.UpdatePostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func (s *postService) Delete(id, userID uint) error {
	post, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if post.UserID != userID {
		return errors.New("unauthorized: you can only delete your own posts")
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
		response[i] = dto.GetPostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
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
		response[i] = dto.GetPostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
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