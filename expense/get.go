package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func GetByIdHandler(c echo.Context) error {
	id := c.Param("id")

	rowID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "id should be int " + err.Error()})
	}

	row := db.QueryRow("SELECT id, title, amount, note, tags FROM expenses WHERE id=$1", rowID)

	exp := Expense{}
	err = row.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, exp)
}

func GetHandler(c echo.Context) error {
	exps := []Expense{}

	rows, err := db.Query("SELECT id, title, amount, note, tags FROM expenses")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	for rows.Next() {
		exp := Expense{}
		err := rows.Scan(&exp.ID, &exp.Title, &exp.Amount, &exp.Note, pq.Array(&exp.Tags))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})

		}
		exps = append(exps, exp)
	}

	return c.JSON(http.StatusOK, exps)
}
