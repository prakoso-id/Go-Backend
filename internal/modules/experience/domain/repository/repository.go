package repository

import (
	"go-backend/internal/modules/experience/domain/entity"
	"gorm.io/gorm"
)

type ExperienceRepository interface {
	Create(experience *entity.Experience) error
	GetAll() ([]entity.Experience, error)
	GetByID(id uint) (*entity.Experience, error)
	GetByUserID(userID uint) ([]entity.Experience, error)
	Update(experience *entity.Experience) error
	Delete(id uint) error
}

type experienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceRepository{
		db: db,
	}
}

func (r *experienceRepository) Create(experience *entity.Experience) error {
	result := r.db.Create(experience)
	return result.Error
}

func (r *experienceRepository) GetAll() ([]entity.Experience, error) {
	var experiences []entity.Experience
	result := r.db.Find(&experiences)
	return experiences, result.Error
}

func (r *experienceRepository) GetByID(id uint) (*entity.Experience, error) {
	var experience entity.Experience
	result := r.db.First(&experience, id)
	return &experience, result.Error
}

func (r *experienceRepository) GetByUserID(userID uint) ([]entity.Experience, error) {
	var experiences []entity.Experience
	result := r.db.Where("user_id = ?", userID).Find(&experiences)
	return experiences, result.Error
}



func (r *experienceRepository) Update(experience *entity.Experience) error {
	result := r.db.Save(experience)
	return result.Error
}

func (r *experienceRepository) Delete(id uint) error {
	result := r.db.Delete(&entity.Experience{}, id)
	return result.Error
}
