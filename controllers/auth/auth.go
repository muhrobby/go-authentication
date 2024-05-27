package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrobby/go-authentication/database"
	"github.com/muhrobby/go-authentication/models/entity"
	"github.com/muhrobby/go-authentication/utils"
)

func Register(c *fiber.Ctx) error {

	var Register entity.Register

	err := c.BodyParser(&Register)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.ErrInternalServerError,
			"message": "failed to parse register",
			"error":   err.Error(),
		})
	}

	if Register.Password != Register.ConfirmPassword {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.ErrInternalServerError,
			"message": "Password not match",
		})
	}

	hash, errHash := entity.HashPassword(Register.Password)

	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.ErrInternalServerError,
			"message": "failed to hash password",
			"error":   err.Error(),
		})
	}

	var User entity.User

	User.Name = Register.Name
	User.Email = Register.Email
	User.Password = hash
	User.RoleID = Register.Role

	db := database.ConnectDB()

	err = db.Create(&User).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.ErrInternalServerError,
			"message": "Failed to Register",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Registered successfully",
		"data":    Register,
	})

}

func Login(c *fiber.Ctx) error {
	var Login entity.Login

	err := c.BodyParser(&Login)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.ErrInternalServerError,
			"message": "Failed to Body parse Login",
			"error":   err.Error(),
		})
	}

	var User entity.User
	db := database.ConnectDB()

	err = db.First(&User, "email =? ", Login.Email).Error

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Wrong credentials",
			"error":   err.Error(),
		})
	}

	err = entity.CheckPasswordHash(User.Password, Login.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Wrong Password",
			"error":   err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"name":  User.Name,
		"email": User.Email,
		"exp":   time.Now().Add(time.Minute * 3).Unix(),
	}

	token, errGenerate := utils.GenerateToken(&claims)

	if errGenerate != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Wrong credentials",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "X-Auth-Token",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 1),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "Login successfully",
		"token":   token,
	})
}
