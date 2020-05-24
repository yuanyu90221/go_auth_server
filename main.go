package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// read PORT from .env
	PORT := os.Getenv("PORT")
	// setup Default Router
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
