package controller

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Burak-Atas/kahve_fali/models"
	"github.com/Burak-Atas/kahve_fali/openai"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Fortunetelling(model *openai.OpenAI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		uid := c.GetString("uid")

		// Görsel dosyasını alıyoruz
		image, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Resim dosyası yüklenemedi"})
			return
		}

		// Benzersiz bir dosya adı oluşturuyoruz
		uniqueID := uuid.New().String()
		extension := filepath.Ext(image.Filename) // Dosyanın orijinal uzantısını alıyoruz
		fileName := uniqueID + extension          // Benzersiz ID ile uzantıyı birleştiriyoruz
		filePath := "./uploads/" + fileName       // Dosya yolunu belirliyoruz

		// Dosyayı belirtilen klasöre kaydediyoruz
		if err := c.SaveUploadedFile(image, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Dosya kaydedilirken bir hata oluştu"})
			return
		}

		newFilePath := "http://localhost:8080/uploads/" + fileName
		fmt.Println("new file path", newFilePath)

		// msg, err := model.NewChat(newFilePath)
		// if err!=nil{
		// 	return
		// }

		var fortuneTellig models.FortuneTelling
		fortuneTellig.AiComment = "msg"
		fortuneTellig.CreatedAt = time.Now()
		fortuneTellig.FortuneID = uniqueID
		fortuneTellig.ImageUrl = newFilePath

		filter := bson.D{primitive.E{Key: "user_id", Value: uid}}
		update := bson.D{{Key: "$push", Value: primitive.E{Key: "fortune_tellings", Value: fortuneTellig}}}
		_, errUpdate := UserCollection.UpdateOne(ctx, filter, update)
		if errUpdate != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "oluşturuldu", "path": filePath})
	}
}
