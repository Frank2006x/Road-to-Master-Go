package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/Frank2006x/Fibre/src/db"
	"github.com/Frank2006x/Fibre/src/model"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var jstSecretKey = []byte(os.Getenv("JWT_SECRET"))

func RegisterUser(c fiber.Ctx) error {
	type request struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var body request;

	
	
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	fmt.Println(body)
	var existingUser model.User
	err:=db.GetCollection("users").FindOne(c.Context(),bson.M{"email": body.Email}).Decode(&existingUser)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists",
		})
	}
	
	hashedPassword,_ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	userDoc:= bson.M{
		"email": body.Email,
		"password": string(hashedPassword),
	}
	_,err=db.GetCollection("users").InsertOne(c.Context(),userDoc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})


}

func LoginUser(c fiber.Ctx) error{
	type request struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	var body request;

	if err:=c.Bind().Body(&body) ; err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	var user model.User ;
	err:=db.GetCollection("users").FindOne(c.Context(),bson.M{"email": body.Email}).Decode(&user); 
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userId": user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	t,err:=token.SignedString(jstSecretKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name: "jwt",
		Value: t,
		HTTPOnly: true,
		Secure: true,
		SameSite: "Strict",
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}


func LogoutUser(c fiber.Ctx) error{
	c.Cookie(&fiber.Cookie{
		Name: "jwt",
		Value: "",
		HTTPOnly: true,
		Secure: true,
		SameSite: "Strict",
		Expires: time.Now().Add(-24 * time.Hour),
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}


	