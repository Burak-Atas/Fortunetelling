package main

import (
	"os"

	"github.com/Burak-Atas/kahve_fali/controller"
	"github.com/Burak-Atas/kahve_fali/openai"
	"github.com/gin-gonic/gin"
)

func main() {

	token := os.Getenv("token")

	model := openai.NewOpenAI(token)
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
