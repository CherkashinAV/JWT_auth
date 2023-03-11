package main

import (
	"github.com/CherkashinAV/finance_app/controllers"
	"github.com/CherkashinAV/finance_app/initializers"
	"github.com/CherkashinAV/finance_app/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.InitEnv()
	initializers.ConnectDb()
}

func main() {
	Router := gin.Default()
	user := Router.Group("/user") 
	{
		user.POST("/signup", controllers.Signup)
		user.POST("/login", controllers.Login)
		user.GET("/is_auth", middlewares.RequireAuth, controllers.CheckIsAuth)
		user.PUT("/update", middlewares.RequireAuth, controllers.UpdateUserInfo)
		user.GET("/get", middlewares.RequireAuth, controllers.GetUser)
	}
	Router.Run()
}