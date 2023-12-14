package main

import (
	"strconv"

	"gofr.dev/pkg/gofr"
)

type Expense struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
}

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {

		return "Hello World!", nil
	})

	//post a new expense
	app.POST("/addExpense", func(ctx *gofr.Context) (interface{}, error) {

		// Read JSON data from the request body
		var requestPayload struct {
			Description string  `json:"description"`
			Amount      float64 `json:"amount"`
		}

		// Insert data into the 'expenses' table
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO expenses (description, amount) VALUES (?, ?)", requestPayload.Description, requestPayload.Amount)

		return "added successfully", err
	})

	//get all expenses
	app.GET("/allExpenses", func(ctx *gofr.Context) (interface{}, error) {
		var expenses []Expense

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM expenses")
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var expense Expense
			if err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount); err != nil {
				return nil, err
			}

			expenses = append(expenses, expense)
		}

		return expenses, nil
	})

	app.PUT("/putExpense/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract the expense ID from the URL parameters
		expenseID := ctx.PathParam("id")

		// Parse the expenseID to an integer (assuming it's an integer)
		id, err := strconv.Atoi(expenseID)
		if err != nil {
			return nil, err
		}

		// Fetch the updated expense data from the request body (assuming JSON)
		var updatedExpense Expense

		// Perform the update in the database
		_, err = ctx.DB().ExecContext(ctx, "UPDATE expenses SET description=?, amount=? WHERE id=?", updatedExpense.Description, updatedExpense.Amount, id)
		if err != nil {
			return nil, err
		}

		// Return a success message or updated expense data
		return "Expense updated successfully", nil
	})

	app.DELETE("/delExpense/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Extract expense ID from the URL parameters
		id := ctx.PathParam("id")

		// Execute the DELETE query
		_, err := ctx.DB().ExecContext(ctx, "DELETE FROM expenses WHERE id=?", id)
		if err != nil {
			return nil, err
		}

		return "Expense deleted successfully", nil
	})

	app.GET("/totalAmount", func(ctx *gofr.Context) (interface{}, error) {
		var totalAmount float64

		rows, err := ctx.DB().QueryContext(ctx, "SELECT amount FROM expenses")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var amount float64
			if err := rows.Scan(&amount); err != nil {
				return nil, err
			}
			totalAmount += amount
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		return totalAmount, nil
	})

	app.Start()
}
