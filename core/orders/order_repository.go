package orders

type OrderRepository interface {
	Save(order Order) error
	FindAll() ([]Order, error)
	FindByID(id uint) (*Order, error)
}