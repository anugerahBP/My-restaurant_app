package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MenuItem struct {
	Name      string             `json:"name"`
	BasePrice float64            `json:"base_price"`
	Toppings  map[string]float64 `json:"toppings"`
	Fillings  map[string]float64 `json:"fillings"`
}

type OrderItem struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Topping string  `json:"topping"`
	Filling string  `json:"filling"`
}

var menuItems = []MenuItem{
	{
		Name:      "pizza",
		BasePrice: 50000,
		Toppings:  map[string]float64{"cheese": 12000, "chicken": 18000, "pepper": 8000},
		Fillings:  map[string]float64{"cheese": 12000, "tomato": 9000, "tuna": 20000},
	},
	{
		Name:      "doughnut",
		BasePrice: 20000,
		Toppings:  map[string]float64{"cheese": 12000, "blueberry": 12000, "sugar glaze": 10000},
		Fillings:  map[string]float64{"blueberry": 12000, "milk cream": 10000, "apple slices": 14000},
	},
	{
		Name:      "pie",
		BasePrice: 45000,
		Toppings:  map[string]float64{"blueberry": 12000, "apple slices": 14000, "pepper": 8000},
		Fillings:  map[string]float64{"cheese": 12000, "chicken": 18000, "tuna": 20000},
	},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/menu", getMenuHandler).Methods("GET")
	router.HandleFunc("/order", orderHandler).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func getMenuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuItems)
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	var order OrderItem
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order.Price = calculateOrderPrice(order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func calculateOrderPrice(order OrderItem) float64 {
	for _, menuItem := range menuItems {
		if menuItem.Name == order.Name {
			orderPrice := menuItem.BasePrice
			if order.Topping != "" {
				toppingPrice, exists := menuItem.Toppings[order.Topping]
				if exists {
					orderPrice += toppingPrice
				}
			}
			if order.Filling != "" {
				fillingPrice, exists := menuItem.Fillings[order.Filling]
				if exists {
					orderPrice += fillingPrice
				}
			}
			return orderPrice
		}
	}
	return 0
}
