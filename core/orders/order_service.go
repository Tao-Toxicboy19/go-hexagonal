package orders

import "errors"

type OrderService interface {
	CreateOrder(order Order) error
	FindAllOrder() ([]Order, error)
	FindOrderById(id uint) (*Order, error)
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

func (s *orderServiceImpl) FindAllOrder() ([]Order, error) {
	orders, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderServiceImpl) FindOrderById(id uint) (*Order, error) {
	if id == 0 {
		return nil, errors.New("invalid ID, must be greater than 0")
	}

	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
