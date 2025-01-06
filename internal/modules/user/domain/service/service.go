package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go-backend/internal/modules/user/domain/entity"
	"go-backend/internal/modules/user/domain/repository"
	"go-backend/internal/modules/user/dto"
)

type UserService interface {
	Create(req *dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(id uint) (*dto.UserResponse, error)
	GetByEmail(email string) (*dto.UserResponse, error)
	List(page, limit int) ([]*dto.UserResponse, error)
	Update(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
	Delete(id uint) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(userID uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetByEmail(email string) (*dto.UserResponse, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) List(page, limit int) ([]*dto.UserResponse, error) {
	users, err := s.repo.List(page, limit)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.UserResponse, len(users))
	for i, user := range users {
		response[i] = &dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return response, nil
}

func (s *userService) Update(id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	if req.Password != "" {
		user.Password = req.Password
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(req.Password); err != nil {
		return nil, err
	}

	// Generate JWT token
	expiresAt := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expiresAt.Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// Update user with token
	user.SetToken(tokenString, expiresAt)
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token:  tokenString,
		UserID: user.ID,
		Status: "success",
	}, nil
}

func (s *userService) Logout(userID uint) error {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}

	user.ClearToken()
	return s.repo.Update(user)
}