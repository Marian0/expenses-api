package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	common "gitlab.com/marian0/expenses-api/common"
	models "gitlab.com/marian0/expenses-api/models"
)

//GetExpenses : Gets all expenses
func GetExpenses(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	expenses := []models.Expense{}
	rows, err := db.Query("select * from expenses")
	common.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		err := rows.Scan(&expense.ID, &expense.PaidAt, &expense.Title, &expense.Amount)
		common.LogFatal(err)

		expenses = append(expenses, expense)
	}

	json.NewEncoder(w).Encode(expenses)
}

// //GetExpense : Get an expense based on id integer parameter
// func GetExpense(w http.ResponseWriter, r *http.Request) {
// 	var expense models.Expense
// 	params := mux.Vars(r)

// 	id, err := strconv.Atoi(params["id"])
// 	common.LogFatal(err)

// 	row := main.DB.QueryRow("select * from expenses where id=$1", id)

// 	err = row.Scan(&expense.ID, &expense.PaidAt, &expense.Title, &expense.Amount)
// 	common.LogFatal(err)

// 	json.NewEncoder(w).Encode(expense)
// }
