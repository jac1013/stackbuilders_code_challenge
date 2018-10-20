package persistence

import . "stackbuilders_pizza/business/models"

func (memory OnMemoryPersistence) translate(order Order) OrderOnMemoryModel {
	return OrderOnMemoryModel{ID: order.ID, Name: order.Name, Address: order.Address, Phone: order.Phone, Size: order.Size, Ingredients: order.Ingredients, Total: order.Total}
}

func (memory OnMemoryPersistence) translateToBusiness(order OrderOnMemoryModel) Order {
	return Order{ID: order.ID, Customer: Customer{Name: order.Name, Address: order.Address, Phone: order.Phone}, Size: order.Size, Ingredients: order.Ingredients, Total: order.Total}
}

func (memory OnMemoryPersistence) translateAllToBusiness(orders []OrderOnMemoryModel) Orders {
	businessOrders := make(Orders, 0)
	for _, order := range orders  {
		businessOrders = append(businessOrders, memory.translateToBusiness(order))
	}
	return businessOrders
}
