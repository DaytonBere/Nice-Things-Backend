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

	route.POST("/api/user/signUp", middleware.RequireAuth, controllers.SignUp)
	route.POST("/api/user/signIn", controllers.SignIn)
	route.POST("/api/user/signOut", controllers.SignOut)
	route.GET("/api/user/getUsers", middleware.RequireAuth, controllers.GetUsers)
	route.GET("/api/user/getUsersTest", controllers.GetUsersTest)
	route.PATCH("/api/user/changePassword", middleware.RequireAuth, controllers.ChangePassword)
	route.GET("/api/user/validate", middleware.RequireAuth, controllers.Validate)

	route.POST("/api/niceThings/createNiceThing",  middleware.RequireAuth, controllers.CreateNiceThing)
	route.PATCH("/api/niceThings/editNiceThing",  middleware.RequireAuth, controllers.EditNiceThing)

	route.Run()
}