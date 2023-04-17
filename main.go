package main

import (
	"Nice-Things-Backend/controllers"
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	route := gin.Default()
	route.Use(middleware.CORSMiddleware())

	route.POST("/api/user/signUp", controllers.SignUp)
	route.POST("/api/user/signIn", controllers.SignIn)
	route.PATCH("/api/user/changePassword", controllers.ChangePassword)

	route.POST("/api/niceThings/getUsers", controllers.GetUsers)
	route.POST("/api/niceThings/createNiceThing",  controllers.CreateNiceThing)
	route.POST("/api/niceThings/getUserNiceThings",  controllers.GetUserNiceThings)

	route.Run()
}