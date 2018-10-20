package persistence

import (
	. "stackbuilders_pizza/business/models"
	"stackbuilders_pizza/business/persistence"
)

var id = 0

type OnMemoryPersistence struct {
	orders []OrderOnMemoryModel
}

func NewOnMemoryPersistence() persistence.OrderPersistence {
	return OnMemoryPersistence{orders: make([]OrderOnMemoryModel, 0)}
}

func (memory OnMemoryPersistence) Save(order Order) (*Order, error) {
	id++
	order.ID = id
	memory.orders = append(memory.orders, memory.translate(order))
	return &order, nil
}

func (memory OnMemoryPersistence) GetAll() Orders {
	return memory.translateAllToBusiness(memory.orders)
}

func (memory OnMemoryPersistence) Get(id int) *Order {
	for _, order := range memory.orders {
		if order.ID == id {
			model := memory.translateToBusiness(order)
			return &model
		}
	}
	return nil
}

type OrderOnMemoryModel struct {
	ID          int
	Name        string
	Address     string
	Phone       string
	Size        string
	Ingredients []string
	Total       int
}
