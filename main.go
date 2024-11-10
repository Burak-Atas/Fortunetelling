package main

import (
	"github.com/Burak-Atas/kahve_fali/controller"
	"github.com/Burak-Atas/kahve_fali/middleware"
	"github.com/Burak-Atas/kahve_fali/openai"
	"github.com/gin-gonic/gin"
)

func main() {

	model := openai.NewOpenAI("")
	model.NewChat("")
	routers := gin.New(func(e *gin.Engine) {})

	routers.Use(gin.Logger())

	routers.POST("/login", controller.SignIn())
	routers.POST("/signup", controller.SignUp())

	routers.Use(middleware.AuthMiddleware())

	groupV1 := routers.Group("/v1")
	groupV1.POST("/fortunetelling", controller.Fortunetelling(model))

	routers.Run()

}
