package auth

import "gopkg.in/guregu/null.v3"

type paramLogin struct {
	Username null.String `json:"username"`
	Password null.String `json:"password"`
}
