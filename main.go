package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	md "github.com/yuanyu90221/go_auth_server/middle"
)

func main() {
	// read PORT from .env
	PORT := os.Getenv("PORT")
	// setup Default Router
	router := gin.Default()
	router.Use(md.Logger())
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.GET("/", func(c *gin.Context) {
		// time.Sleep(5 * time.Second)
		middle := c.MustGet("middle").(string)
		c.JSON(http.StatusOK, gin.H{
			"message": middle,
		})
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// setup quit channel to receive system shutdown
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
