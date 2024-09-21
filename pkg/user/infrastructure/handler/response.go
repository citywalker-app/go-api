package userhandler

import userdomain "github.com/citywalker-app/go-api/pkg/user/domain"

type Response struct {
	JWT         string          `json:"jwt"`
	ConfirmCode string          `json:"confirmCode"`
	User        userdomain.User `json:"user"`
}
