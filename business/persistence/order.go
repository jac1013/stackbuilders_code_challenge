package persistence

import . "stackbuilders_pizza/business/models"

type OrderPersistence interface {
	Save(order Order) (*Order, error)
	GetAll() Orders
	Get(id int) *Order
}
