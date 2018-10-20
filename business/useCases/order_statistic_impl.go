package useCases

import (
	. "stackbuilders_pizza/business/persistence"
	. "stackbuilders_pizza/business/models"
)

type OrderStatisticImpl struct {
	persister OrderPersistence
}

func NewOrderStatistic(persister OrderPersistence) OrderStatisticImpl {
	return OrderStatisticImpl{persister:persister}
}

func (orderStatistic OrderStatisticImpl) Compute() (*Statistic, error) {
	orders := orderStatistic.persister.GetAll()
	count := make(map[string]int, 0)
	averageOrderTotal := 0
	mostValuableCustomer := make(map[string]int, 0)

	for _, order := range orders {
		for _, ingredient := range order.Ingredients {
			count[ingredient] = count[ingredient] + 1
		}

		mostValuableCustomer[order.Name] = mostValuableCustomer[order.Name] + order.Total

		averageOrderTotal += order.Total
	}

	averageOrderTotal = averageOrderTotal / len(orders)

	return nil, nil
}
