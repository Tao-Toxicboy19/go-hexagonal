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

func (r *GormOrderRepository) FindAll() ([]orders.Order, error) {
	var orders []orders.Order
	if result := r.db.Find(&orders); result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (r *GormOrderRepository)  FindByID(id uint) (*orders.Order, error) {
	var order orders.Order
	if result := r.db.First(&order, id); result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}