package orders

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
}

type orderServiceImpl struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Price <= 0 && order.Stock <= 0 {
		return errors.New("total must be greater than")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}
	
	return nil
}
