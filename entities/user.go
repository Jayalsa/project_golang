package entities

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Name               string             `json:"name" bson:"name" binding:"required"`
	Email              string             `json:"email" bson:"email" binding:"required"`
	Password           string             `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm    string             `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Role               string             `json:"role" bson:"role"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordAt    time.Time          `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	Verified           bool               `json:"verified" bson:"verified"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}

type SignupResponse struct {
	Name      string    `json:"name" bson:"name" binding:"required"`
	Email     string    `json:"email" bson:"email" binding:"required"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}
