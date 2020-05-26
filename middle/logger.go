package middle

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("middle", "Logger")
		// before router
		c.Next()
		// after router
		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()
		log.Println("Status:", status)
	}
}
