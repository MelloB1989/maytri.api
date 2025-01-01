package messages

import (
	"errors"
	"maytri/internal/helpers/chat"
	"maytri/models"

	"github.com/MelloB1989/karma/orm"
	"github.com/gofiber/fiber/v2"
)

type ChatRequest struct {
	UserMessage string `json:"user_message"`
}

func PrivateChat(c *fiber.Ctx) error {
	uid := c.Locals("uid").(string)
	name := c.Locals("name").(string)
	chat_id := c.Params("chat_id")
	req := new(ChatRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := chat.PrivateChat(req.UserMessage, chat_id, uid, name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Message sent successfully, wait for Maytri's response",
	})
}

func GetPrivateChatMessages(c *fiber.Ctx) error {
	chat_id := c.Params("chat_id")
	messagesORM := orm.Load(&models.Messages{})
	m, err := messagesORM.QueryRaw(`SELECT * FROM messages
		WHERE chat_id = $1
		ORDER BY created_at DESC;`, chat_id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get messages",
			"err":   err,
		})
	}
	messages, ok := m.([]*models.Messages)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to get messages",
			"err":   errors.New("failed to convert messages to []models.Messages"),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"messages": messages,
	})
}
