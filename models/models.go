package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	TableName    string `karma_table:"users"`
	Id           string `json:"id"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	Age          int    `json:"age"`
	Location     string `json:"location"`
}

type Maytri struct {
	TableName  string    `karma_table:"maytri"`
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	Descrption string    `json:"description"`
	Image      string    `json:"image"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Profession string    `json:"profession"`
	CreatedAt  time.Time `json:"created_at"`
}

type MytriRequest struct {
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Descrption string `json:"description"`
	Profession string `json:"profession"`
}

type Chats struct {
	TableName string    `karma_table:"chats"`
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type ChatParticipants struct {
	TableName     string `karma_table:"chat_participants"`
	Id            string `json:"id"`
	ChatId        string `json:"chat_id"`
	ParticipantId string `json:"participant_id"`
}

type GroupDetails struct {
	TableName       string `karma_table:"group_details"`
	ChatId          string `json:"chat_id"`
	GroupName       string `json:"group_name"`
	GroupImage      string `json:"group_image"`
	GroupDescrption string `json:"group_description"`
}

type Messages struct {
	TableName string    `karma_table:"messages"`
	Id        string    `json:"id"`
	ChatId    string    `json:"chat_id"`
	SenderId  string    `json:"sender_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Claims struct {
	Email string `json:"email"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
	Uid   string `json:"uid"`
	jwt.StandardClaims
}
