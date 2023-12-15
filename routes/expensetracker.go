package main

import (
	"encoding/json"

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

	type RequestPayload struct {
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
	}

	//post a new expense
	app.POST("/addExpense", func(ctx *gofr.Context) (interface{}, error) {

		// Read JSON data from the request body
		var requestPayload RequestPayload

		if err := json.NewDecoder(ctx.Request().Body).Decode(&requestPayload); err != nil {
			return nil, err
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
		var requestPayload RequestPayload
		if err := json.NewDecoder(ctx.Request().Body).Decode(&requestPayload); err != nil {
			return nil, err
		}

		// Fetch the updated expense data from the request body (assuming JSON)
		expenseID := ctx.Param("id")

		// Update data in the 'expenses' table
		_, err := ctx.DB().ExecContext(ctx, "UPDATE expenses SET description = ?, amount = ? WHERE id = ?", requestPayload.Description, requestPayload.Amount, expenseID)

		// Return a success message or updated expense data
		return "Expense updated successfully", err
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
