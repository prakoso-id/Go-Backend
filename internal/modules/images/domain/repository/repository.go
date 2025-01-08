package repository

import (
	"go-backend/internal/modules/images/domain/entity"

	"gorm.io/gorm"
)

type ImagesRepository interface {
	Create(image *entity.Images) error
	GetByID(id uint) (*entity.Images, error)
	Update(image *entity.Images) error
	Delete(id uint) error
	GetByPostID(postID uint) ([]entity.Images, error)
}

type imagesRepository struct {
	db *gorm.DB
}

func NewImagesRepository(db *gorm.DB) ImagesRepository {
	return &imagesRepository{db: db}
}

func (r *imagesRepository) Create(image *entity.Images) error {
	return r.db.Create(image).Error
}

func (r *imagesRepository) GetByID(id uint) (*entity.Images, error) {
	var image entity.Images
	if err := r.db.First(&image, id).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (r *imagesRepository) Update(image *entity.Images) error {
	return r.db.Save(image).Error
}

func (r *imagesRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Images{}, id).Error
}

func (r *imagesRepository) GetByPostID(postID uint) ([]entity.Images, error) {
	var images []entity.Images
	if err := r.db.Where("post_id = ?", postID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}
