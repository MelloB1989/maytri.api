package auth

import (
	"fmt"
	"maytri/models"

	a "github.com/MelloB1989/karma/auth"
	"github.com/MelloB1989/karma/orm"
)

func GetUser(phone string) (a.AuthUserPhone, error) {
	// Karma ORM
	usersORM := orm.Load(&models.Users{})
	userByPhone, err := usersORM.GetByFieldEquals("Phone", phone)
	if err != nil {
		return nil, err
	}
	user, ok := userByPhone.([]*models.Users)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	au := a.NewAuthUserPhone(user[0].Phone, "", user[0].Id)
	au.SetAdditionalClaims(map[string]interface{}{
		"college": user[0].College,
		"branch":  user[0].Branch,
		"year":    user[0].Year,
		"roll":    user[0].Roll,
	})
	return au, nil
}
