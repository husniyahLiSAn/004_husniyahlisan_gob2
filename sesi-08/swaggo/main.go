package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "swaggo/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Order represents the model for an order
type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`
}

var orders []Order
var prevOrderID = 0

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  soberkoder@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
func main() {
	router := mux.NewRouter()
	// Read
	router.HandleFunc("/orders", getOrders).Methods("GET")
	// ReadById
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	// Create
	router.HandleFunc("/orders", createOrders).Methods("POST")
	// Update
	router.HandleFunc("/orders/{orderId}", UpdateOrder).Methods("PUT")
	// Delete
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// CreateOrders godoc
// @Summary Create a new order
// @Description Create a new order with theb input paylod
// @Tags orders
// @Accept json
// @Produce json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrderById godoc
// @Summary Get details of order by id
// @Description Get details of order by id
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "ID"
// @Success 200 {object} Order
// @Router /orders/{orderId} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for _, order := range orders {
		if order.OrderID == inputOrderId {
			json.NewEncoder(w).Encode(orders)
			return
		}
	}

}

// DeleteOrder godoc
// @Summary Delete data order where orderId
// @Description Delete data order where orderId
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "ID"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

}

// UpdateOrder godoc
// @Summary Update data order where orderId
// @Description Update data order where orderId
// @Tags orders
// @Accept json
// @Produce json
// @Param orderId path int true "ID"
// @Success 200 {object} Order
// @Router /orders/{orderId} [put]
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderId := params["orderId"]

	for i, order := range orders {
		if order.OrderID == inputOrderId {
			orders = append(orders[:i], orders[i+1:]...)
			var updateOrder Order
			json.NewDecoder(r.Body).Decode(&updateOrder)
			orders = append(orders, updateOrder)
			json.NewEncoder(w).Encode(updateOrder)
			return
		}
	}

}
