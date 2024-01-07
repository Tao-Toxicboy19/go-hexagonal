package orders

type OrderRepository interface {
	Save(order Order) error
}