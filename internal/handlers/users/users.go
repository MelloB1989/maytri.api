package users

import (
	"fmt"
	"maytri/models"

	"github.com/MelloB1989/karma/errors"
	"github.com/MelloB1989/karma/orm"
	"github.com/MelloB1989/karma/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RegisterUser(c *fiber.Ctx) error {
	req := new(models.User)
	phone := c.Locals("phone").(string)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
			"errors":  errors.NewKarmaError().GetError(4001),
		})
	}

	req.Phone = phone
	req.Id = utils.GenerateID()

	usersORM := orm.Load(&models.User{})
	if err := usersORM.Insert(req); err != nil {
		fmt.Println(err)
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
		"name":  req.Name,
	})
	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"token":   t,
	})
}
