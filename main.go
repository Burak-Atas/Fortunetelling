package main

import (
	"github.com/Burak-Atas/kahve_fali/controller"
	"github.com/Burak-Atas/kahve_fali/openai"
	"github.com/gin-gonic/gin"
)

func main() {

	model := openai.NewOpenAI("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")
	model.NewChat("")
	routers := gin.New(func(e *gin.Engine) {})

	routers.Use(gin.Logger())

	routers.POST("/login", controller.SignIn())
	routers.POST("/signup", controller.SignUp())

	// routers.Use(middleware.AuthMiddleware())

	groupV1 := routers.Group("/v1")
	groupV1.POST("/fortunetelling", controller.Fortunetelling(model))
	groupV1.GET("/getfortunetelling", controller.GetFortuneTelling())
	groupV1.DELETE("/deletefortunetelling", controller.DelFortuneTelling())

	routers.Run()

}
