package repository

import (
	"go-backend/internal/modules/post/domain/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
	GetByID(id uint) (*entity.Post, error)
	Update(post *entity.Post) error
	Delete(id uint) error
	List(offset, limit int) ([]entity.Post, error)
	ListByUserID(userID uint, offset, limit int) ([]entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *entity.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetByID(id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.db.Preload("User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(post *entity.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Post{}, id).Error
}

func (r *postRepository) List(offset, limit int) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.Preload("User").Offset(offset).Limit(limit).Find(&posts).Error
	return posts, err
}

func (r *postRepository) ListByUserID(userID uint, offset, limit int) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.db.Preload("User").Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&posts).Error
	return posts, err
}