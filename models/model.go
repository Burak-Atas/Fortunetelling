package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	// kullanıcı bilgileri
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`

	// kalan hak
	Remaining int `json:"remaining"`

	//signed detail
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`

	//oluşturulma tarihleri
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	//fal bilgisi
	FortuneTellings []FortuneTelling `json:"fortune_tellings"`
}

type FortuneTelling struct {
	FortuneID string    `json:"fortune_id"`
	ImageUrl  string    `json:"image_url"`
	AiComment string    `json:"ai_comment"`
	CreatedAt time.Time `json:"created_at"`
}
