package maytri

import (
	"maytri/models"

	"github.com/MelloB1989/karma/orm"
	"github.com/gofiber/fiber/v2"
)

func GetMyMaytris(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)

	maytriORM := orm.Load(&models.Maytri{})
	m, err := maytriORM.GetByFieldEquals("UserId", uid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "No maytris found",
			"err":   err,
		})
	}
	maytris, ok := m.([]*models.Maytri)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "No maytris found",
			"err":   "failed to convert maytris to []models.Maytri",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data":    maytris,
		"message": "Maytris found",
	})
}
