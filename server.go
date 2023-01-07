package main

import (
	"fmt"
	"os"

	"github.com/kenfemus/assessment/expense"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	e := echo.New()
	e.POST("/expenses", expense.CreateHandler)
	e.GET("/expenses/:id", expense.GetByIdHandler)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
