package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	kafkaService := NewKafkaService()

	r.POST("/", func(c *gin.Context) {
		var wish Wish
		if err := c.ShouldBindJSON(&wish); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		wish.CreatedAt = time.Now().UTC()

		go func() {
			kafkaService.sendWishData(&wish)
		}()

		c.JSON(http.StatusCreated, wish)
	})

	r.Run()
}
