package controllers

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers (c *gin.Context) {

	var body struct {
		Sender int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	fmt.Printf("Get Users BODY: %+v\n", body)


	// Look up requested user
	var currentUser models.User
	initializers.DB.First(&currentUser, "id = ?", body.Sender)


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
}

func CreateNiceThing (c *gin.Context) {
	var body struct {
		Receiver int
		Body string
		Sender int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	fmt.Printf("CREATENICETHING BODY: %+v\n", body)


	// Look up requested user
	var currentUser models.User
	initializers.DB.First(&currentUser, "id = ?", body.Sender)


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
}


func GetUserNiceThings (c *gin.Context) {
	var body struct {
		Receiver int
		Sender int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	fmt.Printf("GETUSERNICETHING BODY: %+v\n", body)


	// Look up requested user
	var currentUser models.User
	initializers.DB.First(&currentUser, "id = ?", body.Sender)

	if (!currentUser.Admin) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Not admin",
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
	fmt.Println("Number of nice things found:", len(allNiceThings))

	for _, niceThing := range allNiceThings {
		fmt.Printf("niceThing.Receiver = %v, type = %T\n", niceThing.Receiver, niceThing.Receiver)
		fmt.Printf("body.Receiver = %v, type = %T\n", body.Receiver, body.Receiver)
		fmt.Println()
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
	
	fmt.Println("retNiceThings:", retNiceThings)
	c.JSON(200, retNiceThings)	
}