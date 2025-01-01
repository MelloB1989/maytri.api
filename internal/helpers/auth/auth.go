package auth

import (
	"fmt"
	"maytri/models"

	a "github.com/MelloB1989/karma/auth"
	"github.com/MelloB1989/karma/orm"
)

func GetUser(phone string) (a.AuthUserPhone, error) {
	// Karma ORM
	usersORM := orm.Load(&models.User{})
	userByPhone, err := usersORM.GetByFieldEquals("Phone", phone)
	if err != nil {
		return nil, err
	}
	user, ok := userByPhone.([]*models.User)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	if len(user) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	au := a.NewAuthUserPhone(user[0].Phone, "", user[0].Id)
	au.SetAdditionalClaims(map[string]interface{}{
		"email": user[0].Email,
		"age":   user[0].Age,
		"phone": user[0].Phone,
		"uid":   user[0].Id,
		"name":  user[0].Name,
	})
	return au, nil
}
