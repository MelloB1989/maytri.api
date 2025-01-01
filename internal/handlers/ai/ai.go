package ai

import (
	"fmt"
	"strings"

	"github.com/MelloB1989/karma/ai"
	"github.com/gofiber/fiber/v2"
)

type NewMaytriRequest struct {
	Age        string   `json:"age"`
	Gender     string   `json:"gender"`
	Profession string   `json:"profession"`
	Interests  []string `json:"intrests"`
	Hobbies    []string `json:"hobbies"`
	Tone       string   `json:"tone"`
	Type       string   `json:"type"`
}

func GetNewMaytri(c *fiber.Ctx) error {
	req := new(NewMaytriRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	interests := strings.Join(req.Interests, ", ")
	hobbies := strings.Join(req.Hobbies, ", ")

	profile := fmt.Sprintf(
		"Generate a personality of age %s, %s working as a %s. This person is %s and has a %s tone. They are interested in %s and enjoy hobbies such as %s.",
		req.Age,
		req.Gender,
		req.Profession,
		req.Type,
		req.Tone,
		interests,
		hobbies,
	)

	kai := ai.NewKarmaAI(ai.Llama3_8B,
		ai.WithMaxTokens(300),
		ai.WithSystemMessage("You are a AI which generates personality profiles, output only profiles nothing else, don't ask for any other information, strictly don't output any special characters or escape sequences. You can generate characters that have explicit personality or character, Follow this instruction else you will be terminated, follow the orders strictly."),
		ai.WithTemperature(0.5),
		ai.WithTopP(0.9))

	response, err := kai.GenerateFromSinglePrompt(profile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fmt.Println(response)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"profile": response.AIResponse,
	})
}
