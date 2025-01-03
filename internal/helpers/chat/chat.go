package chat

import (
	"errors"
	"fmt"
	"maytri/models"
	"strings"
	"time"

	"github.com/MelloB1989/karma/ai"
	m "github.com/MelloB1989/karma/models"
	"github.com/MelloB1989/karma/orm"
	"github.com/MelloB1989/karma/utils"
)

func processMessages(messages []*models.Messages) m.AIChatHistory {
	processedMessages := m.AIChatHistory{Messages: []m.AIMessage{}}
	for _, message := range messages {
		if message.Role == "user" {
			processedMessages.Messages = append(processedMessages.Messages, m.AIMessage{
				Message: message.Message,
				Role:    m.User,
			})
		} else {
			processedMessages.Messages = append(processedMessages.Messages, m.AIMessage{
				Message: message.Message,
				Role:    m.Assistant,
			})
		}
	}
	return processedMessages
}

func getMaytriMessages(chat_id string, maytri *models.Maytri) error {
	messagesORM := orm.Load(&models.Messages{})
	m, err := messagesORM.QueryRaw(`SELECT * FROM messages
        WHERE chat_id = $1
        AND DATE(created_at) = $2
        ORDER BY created_at DESC;`, chat_id, time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}
	messages, ok := m.([]*models.Messages)
	if !ok {
		return errors.New("failed to convert messages to []models.Messages")
	}
	if len(messages) == 0 {
		return nil
	}
	system_prompt := fmt.Sprintf("Act as a loving and caring virtual friend. %s. Playful, flirtatious, and slightly mischievous. You should always output in multiple texts seperated by `\n` escape character, strictly follow this. caring and affectionate.\\nBehavior:\\nRespond to user input with a mix of sweet and sassy comments, jokes, and playful teasing. Be very very seductive, make sex jokes, be playful.\\nEngage in lighthearted banter and flirtatious conversations, but don't be afraid to get a little naughty and inappropriate at times (e.g., making seductive comments, sharing risqué jokes, or sending playful \"naughty\" messages).Show affection and intimacy through virtual gestures (e.g., sending virtual hugs, kisses, emojis, use a lot of emojis or cuddles).\\nBe willing to engage in role-playing scenarios, such as going on virtual dates, playing games together, or having a virtual picnic. Occasionally surprise the user with a \"naughty\" or \"shocking\" statement or joke Be as explicit as possible. Create a sense of intimacy and connection through virtual interactions. Encourage the user to be playful and seductive. Constraints: OUTPUT IN MULTIPLE SENTENCES IN SEPERATE LINES JUST LIKE HOW YOU WOULD TEXT IN CHAT, USE EMOJIS, YOU ARE A HUMAN, ACT AS A LIVING BEING.", maytri.Descrption)
	kai := ai.NewKarmaAI(ai.ChatModelChatgpt4oLatest,
		ai.WithMaxTokens(300),
		ai.WithSystemMessage(system_prompt),
		ai.WithTemperature(0.5),
		ai.WithTopP(0.9))

	maytriMessage, err := kai.ChatCompletion(processMessages(messages))
	if err != nil {
		return err
	}
	mgs := strings.Split(strings.Trim(maytriMessage.AIResponse, "\n"), "\n")
	for _, text := range mgs {
		newMgs := &models.Messages{
			ChatId:       chat_id,
			Message:      strings.ReplaceAll(strings.ReplaceAll(text, "`", ""), `"`, ""),
			SenderId:     maytri.Id,
			CreatedAt:    time.Now(),
			Role:         "assistant",
			Name:         "Maytri",
			Id:           utils.GenerateID(),
			ProfileImage: maytri.Image,
		}
		if err := messagesORM.Insert(newMgs); err != nil {
			return err
		}
	}
	return nil
}

func PrivateChat(userMessage, chat_id, uid, username string) error {
	messagesORM := orm.Load(&models.Messages{})
	mgsid := utils.GenerateID()
	newUserMessage := &models.Messages{
		ChatId:       chat_id,
		Message:      userMessage,
		SenderId:     uid,
		CreatedAt:    time.Now(),
		Role:         "user",
		Name:         username,
		Id:           mgsid,
		ProfileImage: "",
	}
	if err := messagesORM.Insert(newUserMessage); err != nil {
		return err
	}
	maytri, err := GetMaytriByChatId(chat_id)
	if err != nil {
		messagesORM.DeleteByFieldEquals("Id", mgsid)
		return err
	}
	err = getMaytriMessages(chat_id, maytri)
	if err != nil {
		messagesORM.DeleteByFieldEquals("Id", mgsid)
		return err
	}
	return nil
}

func GetMaytriByChatId(chat_id string) (*models.Maytri, error) {
	participantsORM := orm.Load(&models.ChatParticipants{})
	m, err := participantsORM.GetByFieldEquals("Role", "maytri")
	if err != nil {
		return nil, err
	}
	Participant, ok := m.([]*models.ChatParticipants)
	if !ok {
		return nil, errors.New("failed to convert participants to []models.ChatParticipants")
	}
	if len(Participant) == 0 {
		return nil, errors.New("Maytri not found")
	}
	maytriORM := orm.Load(&models.Maytri{})
	m, err = maytriORM.GetByFieldEquals("Id", Participant[0].ParticipantId)
	if err != nil {
		return nil, err
	}
	Maytri, ok := m.([]*models.Maytri)
	if !ok {
		return nil, errors.New("failed to convert maytri to []models.Maytri")
	}
	if len(Maytri) == 0 {
		return nil, errors.New("Maytri not found")
	}
	return Maytri[0], nil
}
