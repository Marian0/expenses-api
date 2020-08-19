package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"gitlab.com/marian0/expenses-api/common"
	"gitlab.com/marian0/expenses-api/handlers"
	"gitlab.com/marian0/expenses-api/models"
)

var expenses []models.Expense

//DB db instance
var DB *sql.DB
var connectionString string

func init() {
	//Load environment variables
	err := godotenv.Load()
	common.LogFatal(err)

	//Set connection string
	connectionString = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/" + os.Getenv("POSTGRES_DB") + "?sslmode=disable"
	pgURL, err := pq.ParseURL(connectionString)
	common.LogFatal(err)

	//Init db connection
	DB, err = sql.Open("postgres", pgURL)
	common.LogFatal(err)

	err = DB.Ping()
	common.LogFatal(err)
}

func getExpenses(w http.ResponseWriter, r *http.Request) {
	handlers.GetExpenses(w, r, DB)
}

func main() {

	log.Println(connectionString)
	log.Println(DB)

	//API endpoint definition
	router := mux.NewRouter()
	router.HandleFunc("/expenses", getExpenses).Methods("GET")
	// router.HandleFunc("/expenses/{id}", handlers.GetExpense).Methods("GET")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses", createExpense).Methods("POST")
	// router.HandleFunc("/expenses/{id}", removeExpense).Methods("DELETE")

	serverHost := ":" + os.Getenv("APP_PORT")
	log.Println("Init server at " + serverHost)
	log.Fatal(http.ListenAndServe(serverHost, router))
}
