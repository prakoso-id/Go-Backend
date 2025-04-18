package repository

import (
	"go-backend/internal/modules/project/domain/entity"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *entity.Project) error
	GetByID(id uint) (*entity.Project, error)
	GetAll() ([]entity.Project, error)
	Update(project *entity.Project) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]entity.Project, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(project *entity.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) GetByID(id uint) (*entity.Project, error) {
	var project entity.Project
	err := r.db.First(&project, id).Error
	return &project, err
}

func (r *projectRepository) GetAll() ([]entity.Project, error) {
	var projects []entity.Project
	err := r.db.Find(&projects).Error
	return projects, err
}

func (r *projectRepository) Update(project *entity.Project) error {
	return r.db.Save(project).Error
}

func (r *projectRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Project{}, id).Error
}

func (r *projectRepository) GetByUserID(userID uint) ([]entity.Project, error) {
	var projects []entity.Project
	err := r.db.Where("user_id = ?", userID).Find(&projects).Error
	return projects, err
}