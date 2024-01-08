package rfid

import (
	"net/http"

	"github.com/Adan1996/smarthome/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllRfid(c *gin.Context) {
	var rfids []models.Rfid

	models.DB.Find(&rfids)
	c.IndentedJSON(http.StatusOK, gin.H{"rfid": rfids})
}

func GetRfidByUid(c *gin.Context) {
	var rfid models.Rfid
	uid := c.Param("uid")

	if err := models.DB.Where("uid = ?", uid).Find(&rfid).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "RFID not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"rfid": rfid})
}

func PostRfid(c *gin.Context) {
	var rfid models.Rfid

	if err := c.ShouldBindJSON(&rfid); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&rfid)
	c.IndentedJSON(http.StatusOK, gin.H{"rfid": rfid})
}
