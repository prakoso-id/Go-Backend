package repository

import (
	"go-backend/internal/modules/socialmedia/domain/entity"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia *entity.SocialMedia) error
	GetByID(id uint) (*entity.SocialMedia, error)
	GetAll() ([]entity.SocialMedia, error)
	Update(socialMedia *entity.SocialMedia) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]entity.SocialMedia, error)
	GetByProfileID(profileID uint) ([]entity.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) Create(socialMedia *entity.SocialMedia) error {
	return r.db.Create(socialMedia).Error
}

func (r *socialMediaRepository) GetByID(id uint) (*entity.SocialMedia, error) {
	var socialMedia entity.SocialMedia
	err := r.db.Preload("User").First(&socialMedia, id).Error
	return &socialMedia, err
}

func (r *socialMediaRepository) GetAll() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	err := r.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) Update(socialMedia *entity.SocialMedia) error {
	return r.db.Save(socialMedia).Error
}

func (r *socialMediaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.SocialMedia{}, id).Error
}

func (r *socialMediaRepository) GetByUserID(userID uint) ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) GetByProfileID(profileID uint) ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	err := r.db.Preload("User").Where("profile_id = ?", profileID).Find(&socialMedias).Error
	return socialMedias, err
}