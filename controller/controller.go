package controller

import (
	"context"
	"time"

	"github.com/Burak-Atas/kahve_fali/database"
	"github.com/Burak-Atas/kahve_fali/jwt"
	"github.com/Burak-Atas/kahve_fali/models"
	"github.com/Burak-Atas/kahve_fali/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection = database.UserCollection("Users", database.Client)

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.UserModel
		var foundUser models.UserModel

		if err := c.BindJSON(&user); err != nil {
			c.JSON(500, gin.H{
				"error": "",
			})
			return
		}

		filter := bson.D{primitive.E{Key: "email", Value: user.Email}}
		err := UserCollection.FindOne(ctx, filter).Decode(&foundUser)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "",
			})
			return
		}

		ok, msg := utils.VerifyPassword(foundUser.Password, user.Password)
		if !ok {
			c.JSON(400, gin.H{
				"error": msg,
			})
			return
		}

		token, refreshToken, _ := jwt.TokenGenerator(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserId)

		jwt.UpdateAllTokens(token, foundUser.RefreshToken, refreshToken)

		c.JSON(200, gin.H{
			"token":         token,
			"refresh_token": refreshToken,
			"user_id":       foundUser.UserId,
			"email":         foundUser.Email,
		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user models.UserModel
		if err := c.BindJSON(&user); err != nil {
			return
		}

		user.CreatedAt = time.Now().Local()
		user.FortuneTellings = make([]models.FortuneTelling, 0)
		user.Remaining = 3
		user.ID = primitive.NewObjectID()
		user.UserId = user.ID.Hex()

		token, refresh_token, err := jwt.TokenGenerator(user.Email, user.FirstName, user.LastName, user.UserId)
		if err != nil {
			return
		}

		passwordHash := utils.HashPassword(user.Password)
		user.Password = passwordHash
		user.Token = token
		user.RefreshToken = refresh_token

		_, errInsert := UserCollection.InsertOne(ctx, user)
		if errInsert != nil {
			return
		}

		c.JSON(200, gin.H{
			"message":       "kaydınız oluşturuldu",
			"token":         token,
			"refresh_token": refresh_token,
		})
	}
}

func GetKnowlodgeToday() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
