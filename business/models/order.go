package models

type Order struct {
	Customer
	ID         int
	Size       string
	Ingredients Ingredients
	Total      int
}

var sizePrices = make(map[string]int, 0)

func (order Order) GetSizePrice() int {
	sizePrices["Small"] = 5
	sizePrices["Medium"] = 8
	sizePrices["Large"] = 11

	return sizePrices[order.Size]
}


type Orders []Order
