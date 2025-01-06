package repository

import (
	"go-backend/internal/modules/user/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	List(page, limit int) ([]*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *entity.User) error {
	// Only update non-password fields to avoid rehashing
	return r.db.Model(user).Select("name", "email", "token", "expires_at", "updated_at").Updates(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepository) List(page, limit int) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}