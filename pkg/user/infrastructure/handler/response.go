package userhandler

import userdomain "github.com/citywalker-app/go-api/pkg/user/domain"

type Response struct {
	JWT         string          `json:"jwt"`
	User        userdomain.User `json:"user"`
	ConfirmCode string          `json:"confirmCode"`
}
