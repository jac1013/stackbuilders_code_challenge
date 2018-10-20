package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"

	. "stackbuilders_pizza/persistence"
	. "stackbuilders_pizza/business/persistence"
	"encoding/json"
)

import (
	. "stackbuilders_pizza/business/models"
	"strconv"
	"stackbuilders_pizza/business/useCases"
	"log"
)


var memory OrderPersistence

func main() {
	memory = NewOnMemoryPersistence()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.POST("/order", orderCreateHandler)
	r.GET("/orders/:id", findOrderHandler)
	r.GET("/orders", findAllOrdersHandler)
	r.GET("/statistic", calculateStatistic)

	r.Run()
}

func orderCreateHandler(c *gin.Context) {
	var createOrder CreateOrderRestModel
	err := json.NewDecoder(c.Request.Body).Decode(&createOrder)

	handleError(err, c)

	order, _ := useCases.NewOrderOperation(memory).Create(translateRestToBusiness(createOrder))

	orderJSONString, err := translateBusinessToJSONString(*order)

	handleError(err, c)

	c.String(http.StatusOK, orderJSONString)
}

func handleError (err error, c *gin.Context) {
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func findOrderHandler(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	handleError(err, c)

	order := memory.Get(intId)

	orderJSONString, err := translateBusinessToJSONString(*order)

	handleError(err, c)

	c.String(http.StatusOK, orderJSONString)
}

func findAllOrdersHandler(c *gin.Context) {
	orders := memory.GetAll()

	for _, order := range orders {
		log.Println(order.ID)
	}

	ordersJSON := translateAllBusinessToJSON(orders)

	bytes, err := json.Marshal(ordersJSON)

	handleError(err, c)

	c.String(http.StatusOK, string(bytes))
}

func calculateStatistic(c *gin.Context) {
	c.String(http.StatusOK, "statistics")
}


type CreateOrderRestModel struct {
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	Phone       string   `json:"phone"`
	Size        string   `json:"size"`
	Ingredients []string `json:"ingredients"`
}

type OrderJSON struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	Phone       string   `json:"phone"`
	Size        string   `json:"size"`
	Ingredients []string `json:"ingredients"`
	Total       int      `json:"total"`
}

func translateRestToBusiness(order CreateOrderRestModel) Order {
	return Order{Customer: Customer{Name: order.Name, Address: order.Address, Phone: order.Phone}, Size: order.Size, Ingredients: order.Ingredients}
}

func translateBusinessToJSONModel(order Order) OrderJSON {
	return OrderJSON{ID: order.ID, Name: order.Name, Address: order.Address, Phone: order.Phone, Size: order.Size, Ingredients: order.Ingredients , Total: order.Total}
}

func translateBusinessToJSONString(order Order) (string, error) {
	orderJSON:= translateBusinessToJSONModel(order)

	bytes, err := json.Marshal(orderJSON)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func translateAllBusinessToJSON(orders Orders) []OrderJSON {
	jsonOrders := make([]OrderJSON, 0)
	for _, order := range orders {
		orderJSON := translateBusinessToJSONModel(order)
		jsonOrders = append(jsonOrders, orderJSON)
	}

	return jsonOrders
}

