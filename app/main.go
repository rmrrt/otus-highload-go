package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
	City      string `json:"city"`
	Interests string `json:"interests"`
	Sex       string `json:"sex"`
}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.POST("/user/register", func(c *gin.Context) {

		var user User

		if c.ShouldBind(&user) != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "WRONG JSON"})
			log.Print(&user)
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "OKAY"})
	})

	err = r.Run()
	if err != nil {
		return
	}
}
