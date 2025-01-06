package seeder

import (
	"go-backend/internal/modules/user/domain/entity"
	"log"

	"gorm.io/gorm"
)

var adminUsers = []entity.User{
	{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: "admin123",
	},
	{
		Name:     "Super Admin",
		Email:    "superadmin@example.com",
		Password: "superadmin123",
	},
}

// SeedAdminUsers seeds admin users into the database
func SeedAdminUsers(db *gorm.DB) error {
	for _, admin := range adminUsers {
		var existingUser entity.User
		result := db.Where("email = ?", admin.Email).First(&existingUser)
		
		if result.Error == gorm.ErrRecordNotFound {
			if err := db.Create(&admin).Error; err != nil {
				log.Printf("Error seeding admin user %s: %v", admin.Email, err)
				return err
			}
			log.Printf("Admin user seeded successfully: %s", admin.Email)
		} else if result.Error != nil {
			log.Printf("Error checking existing admin user %s: %v", admin.Email, result.Error)
			return result.Error
		} else {
			log.Printf("Admin user already exists: %s", admin.Email)
		}
	}
	return nil
}
