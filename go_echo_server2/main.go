package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type payment struct {
	ID       string  `json:"id"`
	Invoice  string  `json:"invoice"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

var payments = []payment{}

func main() {

	// in case we need to seed payments data.
	payments = []payment{
		{ID: "1", Invoice: "X672762", Currency: "USD", Amount: 100},
		{ID: "2", Invoice: "E35565", Currency: "EUR", Amount: 200},
		{ID: "3", Invoice: "J35565", Currency: "JPN", Amount: 3000},
	}
	// create a new echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthCheck)

	e.GET("/payments", getPayments)
	e.GET("/payments/:id", getPaymentByID)
	e.GET("/c", countPayments)

	e.POST("/payments", postPayment)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Status is OK")
}

func getPayments(c echo.Context) error {
	return c.JSON(http.StatusOK, payments)
}

func countPayments(c echo.Context) error {

	count := fmt.Sprintf("count: %d", len(payments))
	return c.JSON(http.StatusOK, count)
}

func getPaymentByID(c echo.Context) error {

	var p payment

	id := c.Param("id")

	//fmt.Println("id=", c.Param("id"), id)

	for _, p = range payments {
		if p.ID == id {
			//	fmt.Println("p=", p)
			return c.JSON(http.StatusFound, p)
		}
	}
	return c.JSON(http.StatusNotFound, "payment not found")
}

func postPayment(c echo.Context) error {

	var newPayment payment

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&newPayment)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	payments = append(payments, newPayment)
	return c.String(http.StatusOK, "We got your payment!")
}
