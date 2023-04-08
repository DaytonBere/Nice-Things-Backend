package controllers

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNiceThing (c *gin.Context) {
	userInf, exist := c.Get("user")

	var currentUser models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var body struct {
		Receiver int
		Body string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	var niceThing models.NiceThing
	initializers.DB.Where("Sender = ? AND Receiver = ?", currentUser.ID, body.Receiver).Find(&niceThing)

	if niceThing.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Nice Thing already exists between these users",
		})

		return
	}

	newNiceThing := models.NiceThing{
		Sender: int(currentUser.ID),
		Receiver: body.Receiver,
		Body: body.Body,
	}

	result := initializers.DB.Create(&newNiceThing)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create Nice Thing",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func EditNiceThing (c *gin.Context) {
	userInf, exist := c.Get("user")

	var currentUser models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


	var body struct {
		Receiver int
		Body string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	var niceThing models.NiceThing
	initializers.DB.Where("Sender = ? AND Receiver = ?", currentUser.ID, body.Receiver).Find(&niceThing)

	if niceThing.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Nice Thing does not already exist between these users",
		})

		return
	}

	initializers.DB.Model(&niceThing).Update("Body", body.Body)

	c.JSON(200, gin.H{})
}