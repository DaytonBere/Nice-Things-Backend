package controllers

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp (c *gin.Context) {

	userInf, exist := c.Get("user")

	var currentUser models.User = userInf.(models.User)

	if !exist || !currentUser.Admin {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Get email, first name, last name, and if they are an admin

	var body struct {
		Email string
		FirstName string
		LastName string
		Admin bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	// Create password = lower(LastName + FirstName)
	password := strings.ToLower(body.LastName + body.FirstName)

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash a password",
		})

		return
	}

	// Create User
	user := models.User{
		Email: body.Email, 
		FirstName: body.FirstName, 
		LastName: body.LastName, 
		Password: string(hash), 
		Admin: body.Admin,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Default Password": password,
	})
}

func SignIn (c *gin.Context) {
	// Get the email and pass off req body
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid email",
		})

		return
	}

	// Compare submitted password with saved user password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid password",
		})

		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create token",
		})

		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"_id": user.ID,
		"email": user.Email,
		"token": tokenString,
	})
}	

func SignOut (c *gin.Context) {
	_, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func ChangePassword (c *gin.Context) {

	var body struct {
		OldPassword string
		NewPassword string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	userInf, exist := c.Get("user")

	var user models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid password",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash a password",
		})

		return
	}

	initializers.DB.Model(&user).Update("Password", string(hash))

	c.JSON(200, gin.H{})
}

func Validate (c *gin.Context) {
	userInf, exist := c.Get("user")

	var user models.User = userInf.(models.User)

	if !exist {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "I am logged in",
		"user": user,
	})
}

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
}

func GetUsersTest (c *gin.Context) {
	var allUsers []models.User
	initializers.DB.Find(&allUsers)

	c.JSON(http.StatusOK, allUsers)
}