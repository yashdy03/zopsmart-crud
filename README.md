# BudgetBuddy | GOFR

An expense tracker for you to keep track of your expenditure and be resourceful of your budget built using GOFR "An Opinionated Go Framework"

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Postman Screenshots](#postman-screenshots)

## Getting Started

Welcome to BudgetBuddy, your go-to application for managing and tracking your expenses effortlessly.


### Prerequisites

- [Go](https://golang.org/) (Programming Language)
- [GoFr](https://gofr.dev/) (Framework)
- [MySQL](https://www.mysql.com/) (or any preferred database)
- [Git](https://git-scm.com/) (Version Control System)


### Installation

step-by-step instructions on how to install and set up .

- git clone https://github.com/yashdy03/zopsmart-crud
- cd routes
- go run expensetracker.go

## Usage

### View All Expenses:

- Explore the "All Expenses" route to view a list of all your recorded expenses.

### Calculate Total Amount:

- Check the "Total Amount" route to see the overall sum of your expenses.

### Update an Expense:

- Visit the "Update Expense" route to modify details of a specific expense using its id.

## Postman Screenshots

### Add Expense:

1. Open Postman and set up a `POST` request to the "addExpense" endpoint.
2. Provide the {"description":"amount"} in the request body.
3. Send the request.

   ![Add Expense](screenshots/postman_add_expense.png)

### Delete Expense:

1. Open Postman and set up a `DELETE` request to the "deleteExpense/{id}" endpoint.
2. Send the request.

   ![Add Expense](screenshots/postman_add_expense.png)

### View All Expenses:

1. Set up a `GET` request to your "/allExpenses" endpoint.
2. Send the request.

   ![View Expenses](screenshots/postman_view_expenses.png)

### Update an Expense:

1. Configure a `PUT` request to your "/putExpense/{id}" endpoint.
2. Adjust the request body with the updated details.
3. Send the request.

   ![Update Expense](screenshots/postman_update_expense.png)

### Calculate Total Amount:

1. Create a `GET` request to your "/totalAmount" endpoint.
2. Send the request.

   ![Total Amount](screenshots/postman_total_amount.png)



