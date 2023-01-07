package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kenfemus/assessment/expense"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if strings.Contains(token, "wrong_token") {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	})

	e.POST("/expenses", expense.CreateHandler)
	e.GET("/expenses/:id", expense.GetByIdHandler)
	e.PUT("/expenses/:id", expense.UpdateHandler)
	e.GET("/expenses", expense.GetHandler)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
