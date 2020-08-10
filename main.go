package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Expense struct {
	ID     int    `json:id`
	PaidAt string `json:paidAt`
	Title  string `json:title`
	Amount int    `json:amount`
}

var expenses []Expense
var db *sql.DB
var connectionString string

//Check fatal error
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	//Load environment variables
	err := godotenv.Load()
	logFatal(err)

	//Set connection string
	connectionString = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	pgURL, err := pq.ParseURL(connectionString)
	logFatal(err)

	//Init db connection
	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)
}

func main() {

	log.Println(connectionString)
	log.Println(db)

	//API endpoint definition
	router := mux.NewRouter()
	router.HandleFunc("/expenses", getExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", getExpense).Methods("GET")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses/{id}", removeExpense).Methods("DELETE")

	serverHost := ":" + os.Getenv("APP_PORT")
	log.Println("Init server at " + serverHost)
	log.Fatal(http.ListenAndServe(serverHost, router))
}

//Get all expenses
func getExpenses(w http.ResponseWriter, r *http.Request) {
	expenses = []Expense{}
	rows, err := db.Query("select * from expenses")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.ID, &expense.PaidAt, &expense.Title, &expense.Amount)
		logFatal(err)

		expenses = append(expenses, expense)
	}

	json.NewEncoder(w).Encode(expenses)
}

//Get an expense based on id integer parameter
func getExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	logFatal(err)

	row := db.QueryRow("select * from expenses where id=$1", id)

	err = row.Scan(&expense.ID, &expense.PaidAt, &expense.Title, &expense.Amount)
	logFatal(err)

	json.NewEncoder(w).Encode(expense)
}
