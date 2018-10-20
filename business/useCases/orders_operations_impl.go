package useCases

import (
	. "stackbuilders_pizza/business/models"
	. "stackbuilders_pizza/business/persistence"
)

type OrdersOperationImpl struct {
	persister OrderPersistence
}

func NewOrderOperation(persister OrderPersistence) OrdersOperationImpl {
	return OrdersOperationImpl{persister:persister}
}

func (ordersOperation OrdersOperationImpl) Create(order Order) (*Order, error) {
	calculateTotalPrice(&order)
	model, _ := ordersOperation.persister.Save(order)
	return model, nil
}

func calculateTotalPrice(order *Order) {
	total := 0

	for _, ingredient := range order.Ingredients {
		total += GetPrice(ingredient)
	}

	order.Total = total + order.GetSizePrice()
}
