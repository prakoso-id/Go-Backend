package repository

import (
	"go-backend/internal/modules/profile/domain/entity"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(profile *entity.Profile) error
	GetByID(id uint) (*entity.Profile, error)
	GetAll() ([]entity.Profile, error)
	Update(profile *entity.Profile) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]entity.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) Create(profile *entity.Profile) error {
	return r.db.Create(profile).Error
}

func (r *profileRepository) GetByID(id uint) (*entity.Profile, error) {
	var profile entity.Profile
	err := r.db.First(&profile, id).Error
	return &profile, err
}

func (r *profileRepository) GetAll() ([]entity.Profile, error) {
	var profiles []entity.Profile
	err := r.db.Find(&profiles).Error
	return profiles, err
}

func (r *profileRepository) Update(profile *entity.Profile) error {
	return r.db.Save(profile).Error
}

func (r *profileRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Profile{}, id).Error
}

func (r *profileRepository) GetByUserID(userID uint) ([]entity.Profile, error) {
	var profiles []entity.Profile
	err := r.db.Where("user_id = ?", userID).Find(&profiles).Error
	return profiles, err
}