package adapters

import (
	"gorm.io/gorm"
	"toxicboy/core/orders"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) orders.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order orders.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
