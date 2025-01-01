package models

type User struct {
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
