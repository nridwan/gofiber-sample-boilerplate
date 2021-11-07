package auth

import (
	"github.com/nridwan/models"
	"github.com/volatiletech/null/v8"
)

type UserDto struct {
	ID       int64       `json:"id"`
	Username string      `json:"username"`
	Name     null.String `json:"name,omitempty"`
}

func getUserDto(user *models.User) UserDto {
	return UserDto{
		user.ID,
		user.Username,
		user.Name,
	}
}
