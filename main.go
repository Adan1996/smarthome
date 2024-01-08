package main

import (
	"github.com/Adan1996/smarthome/controllers/rfid"
	"github.com/Adan1996/smarthome/models"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()
	models.ConnectionDatabase()
	router.Use(corsMiddleware())

	router.GET("/rfids", rfid.GetAllRfid)
	router.GET("/rfids/:uid", rfid.GetRfidByUid)
	router.POST("/rfids/post", rfid.PostRfid)

	router.Run("localhost:9000")
}
