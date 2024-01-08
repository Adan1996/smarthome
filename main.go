package main

import (
	"github.com/Adan1996/smarthome/controllers/rfid"
	"github.com/Adan1996/smarthome/models"

	"github.com/gin-gonic/gin"
)

// var rfids = []Rfid{
// 	{1, "PICC", "1234567890123456"},
// 	{2, "PICC", "1234567890123456"},
// 	{3, "PICC", "1234567890123456"},
// 	{4, "PICC", "1234567890123456"},
// 	{5, "PICC", "1234567890123456"},
// 	{6, "PICC", "1234567890123456"},
// }

// func getAllRfid(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, rfids)
// }

// func postRfid(c *gin.Context) {
// 	var newRfid Rfid
// 	if err := c.BindJSON(&newRfid); err != nil {
// 		return
// 	}

// 	rfids = append(rfids, newRfid)
// 	c.IndentedJSON(http.StatusCreated, newRfid)
// }

// func getRfidByUid(c *gin.Context) {
// 	id := c.Param("uid")

// 	for _, a := range rfids {
// 		if a.UID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "RFID not found"})
// }

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
