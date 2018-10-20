package useCases

import (
	. "stackbuilders_pizza/business/models"
)

type OrdersOperation interface {
	Create(order Order) (Order, error)
}
