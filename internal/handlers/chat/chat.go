package chat

import (
	"maytri/models"
	"time"

	"github.com/MelloB1989/karma/orm"
	"github.com/MelloB1989/karma/utils"
	"github.com/gofiber/fiber/v2"
)

func CreatePrivateChat(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	req := new(models.MytriRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	//Create Maytri
	mid := utils.GenerateID()
	maytriORM := orm.Load(&models.Maytri{})
	newMaytri := &models.Maytri{
		Id:         mid,
		UserId:     uid,
		Descrption: req.Descrption,
		Image:      "",
		Age:        req.Age,
		Gender:     req.Gender,
		Profession: req.Profession,
		CreatedAt:  time.Now(),
	}
	if err := maytriORM.Insert(newMaytri); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//Create Chat
	cid := utils.GenerateID()
	chatsORM := orm.Load(&models.Chats{})
	newChat := &models.Chats{
		Id:        cid,
		UserId:    uid,
		Type:      "private",
		CreatedAt: time.Now(),
	}
	if err := chatsORM.Insert(newChat); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//Add Maytri to Chat
	participantsORM := orm.Load(&models.ChatParticipants{})
	if err := participantsORM.Insert(&models.ChatParticipants{
		Id:            utils.GenerateID(),
		ChatId:        cid,
		ParticipantId: mid,
	}); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	//Add User to Chat
	if err := participantsORM.Insert(&models.ChatParticipants{
		Id:            utils.GenerateID(),
		ChatId:        cid,
		ParticipantId: uid,
	}); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Chat created successfully",
		"data":    newChat,
	})
}
