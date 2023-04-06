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

	route.POST("/api/user/signUp", controllers.SignUp)
	route.POST("/api/user/signIn", controllers.SignIn)
	route.POST("/api/user/signOut", controllers.SignOut)
	route.PATCH("/api/user/changePassword", middleware.RequireAuth, controllers.ChangePassword)
	route.GET("/api/user/validate", middleware.RequireAuth, controllers.Validate)

	route.Run()
}