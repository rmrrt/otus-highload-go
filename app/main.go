package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", "host=localhost port=5432 user=youruser dbname=yourdb password=yourpassword sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Something"})
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

	r.Run()
}
