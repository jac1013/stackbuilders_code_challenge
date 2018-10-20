package models

type Ingredients []string

var prices = make(map[string]int, 0)


func GetPrice(ingredientName string) int {
  prices["Extra cheese"] = 1
  prices["Anchovies"] = 2
  prices["Pineapple"] = 1
  prices["Onions"] = 1
  prices["Caviar"] = 5
  prices["Kobe Beef"] = 10

  return prices[ingredientName]
}
