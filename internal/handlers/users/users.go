package users

import (
	"maytri/models"

	"github.com/MelloB1989/karma/errors"
	"github.com/MelloB1989/karma/orm"
	"github.com/MelloB1989/karma/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RegisterUser(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
			"errors":  errors.NewKarmaError().GetError(4001),
		})
	}

	usersORM := orm.Load(&models.User{})
	if err := usersORM.Insert(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
			"errors":  errors.NewKarmaError().GetError(5001),
		})
	}
	t, _ := utils.GenerateJWT(jwt.MapClaims{
		"email": req.Email,
		"age":   req.Age,
		"phone": req.Phone,
		"uid":   req.Id,
	})
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"token":   t,
	})
}
