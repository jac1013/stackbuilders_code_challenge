package models

type Order struct {
	Customer
	ID         int
	Size       string
	Ingredients Ingredients
	Total      int
}

type Orders []Order
