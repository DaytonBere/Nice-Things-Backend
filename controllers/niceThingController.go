package controllers

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers (c *gin.Context) {
	userInf, exist := c.Get("user")

	var currentUser models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


	type RetUser struct {
		ID int
		FullName string
		SentNiceThing bool
	}

	users := []RetUser{}

	var allUsers []models.User
	initializers.DB.Find(&allUsers)


	for _, user := range allUsers {
		if user.ID == 1{
			continue
		}

		var niceThing models.NiceThing
		initializers.DB.Where("Sender = ? AND Receiver = ?", currentUser.ID, user.ID).Find(&niceThing)

		retUserInfo := RetUser{
			ID: int(user.ID),
			FullName: user.FirstName + " " + user.LastName,
			SentNiceThing: niceThing.ID != 0,
		}

		users = append(users, retUserInfo)
	}


	c.JSON(http.StatusOK, users)
	return
}

func GetUsersTest (c *gin.Context) {
	var allUsers []models.User
	initializers.DB.Find(&allUsers)

	c.JSON(http.StatusOK, allUsers)
	return
}

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

	if body.Receiver == 1 || currentUser.ID == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Attempting to send to or from super admin",
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
	return
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

	if body.Receiver == 1 || currentUser.ID == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Attempting to send to or from super admin",
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
	return
}

func GetUserNiceThings (c *gin.Context) {

	fmt.Println("HERE")
	userInf, exist := c.Get("user")

	var _ models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var body struct {
		Receiver int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	type RetNiceThing struct {
		ID int
		FullNameSender string
		Body string
	}

	var allNiceThings []models.NiceThing
	retNiceThings := []RetNiceThing{}

	initializers.DB.Find(&allNiceThings)

	for _, niceThing := range allNiceThings {
		if niceThing.Receiver == body.Receiver {
			var user models.User
			initializers.DB.Where("ID = ?", niceThing.Sender).Find(&user)

			toAppend := RetNiceThing{
				ID: int(niceThing.ID),
				FullNameSender: user.FirstName + " " + user.LastName,
				Body: niceThing.Body,
			}

			retNiceThings = append(retNiceThings, toAppend)
		}
	}
	
	c.JSON(200, retNiceThings)	
	return
}