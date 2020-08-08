package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Expense struct {
	ID     int    `json:id`
	Date   string `json:date`
	Title  string `json:title`
	Amount int    `jsom:amount`
}

var expenses []Expense

func main() {

	router := mux.NewRouter()

	expenses = append(expenses,
		Expense{ID: 1, Date: "2020-01-01", Title: "Car repairng", Amount: 15000},
		Expense{ID: 2, Date: "2020-01-02", Title: "Coke", Amount: 500},
		Expense{ID: 3, Date: "2020-01-03", Title: "Gourmet salad", Amount: 10000},
		Expense{ID: 4, Date: "2020-01-04", Title: "Pizza slice", Amount: 150},
	)

	router.HandleFunc("/expenses", getExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", getExpense).Methods("GET")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses/{id}", removeExpense).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}

//Get all expenses
func getExpenses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(expenses)
}

//Get an expense based on id integer parameter
func getExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Try to convert parameter into int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		//@todo: respond a 403 error
		log.Println("Error: Id should be integer")
		return
	}

	for _, expense := range expenses {
		if expense.ID == id {
			json.NewEncoder(w).Encode(&expense)
			return
		}
	}

	//@todo: refactor this to respond an appropiated json response
	json.NewEncoder(w).Encode(`{
		"erro"r: "Not found"
	}`)

}
