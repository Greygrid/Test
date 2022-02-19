package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host     = "localhost"
	port     = 2283
	user     = "postgres"
	password = "cosmoboty"
	dbname   = "exercise"
)

func main() {
	e := echo.New()
	e.GET("/hello", get)
	e.POST("/payments", save)
	e.PUT("/payments", change)
	e.DELETE("/payments", delete)
	e.Logger.Fatal(e.Start(":1323"))

}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "halo")
}

func save(c echo.Context) error {
	price := c.FormValue("price")
	date := c.FormValue("date")
	types := c.FormValue("type")
	name := c.FormValue("name")
	comments := c.FormValue("comments")
	category := c.FormValue("category")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connection successful!")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	sqlStatement := `
	INSERT INTO payments 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(sqlStatement, name, price, date, types, comments, category)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert successful")
	return c.String(http.StatusOK, "price:"+price+", name:"+name+", date:"+date+", type:"+types+", comment:"+comments+", category"+category)
}

func change(c echo.Context) error {
	price := c.FormValue("price")
	date := c.FormValue("date")
	types := c.FormValue("type")
	name := c.FormValue("name")
	comments := c.FormValue("comments")
	category := c.FormValue("category")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
	UPDATE payments
	SET price = $2, date = $3, type = $4, comments =$5, category = $6
	WHERE name = $1;`
	_, err = db.Exec(sqlStatement, name, price, date, types, comments, category)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, "price:"+price+", name:"+name+", date:"+date+", type:"+types+", comment:"+comments+", category"+category)
}

func delete(c echo.Context) error {
	name := c.FormValue("name")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
	DELETE FROM payments
	WHERE name = $1;`
	_, err = db.Exec(sqlStatement, name)
	if err != nil {
		panic(err)
	}
	return c.String(http.StatusOK, "Delete payment`s name:"+name)
}
