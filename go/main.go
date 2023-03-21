package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/uptrace/bun"
)

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Content       string    `bun:"content,notnull"`
	Done          bool      `bun:"done"`
	Until         time.Time `bun:"until,nullzero"`
	CreatedAt     time.Time
	UpdateAt      time.Time `bun:",nullzero"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

func main() {
	fmt.Println("Hello World")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8989"))
}
