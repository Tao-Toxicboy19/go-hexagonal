// adapters/database/gorm_adapter.go

package database

import (
	"gorm.io/gorm"
	"toxicboy/core/users"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) users.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user users.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormUserRepository) FindAll() ([]users.User, error) {
	var users []users.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *GormUserRepository) FindByID(id uint) (*users.User, error) {
	var user users.User
	if result := r.db.First(&user, id); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}