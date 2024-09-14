package userhandler

import userdomain "github.com/citywalker-app/go-api/pkg/user/domain"

type TestCase struct {
	Name       string
	User       userdomain.User
	StatusCode int
}

var userWithNoEmail = userdomain.User{
	Email:    "",
	Password: "testingtesting",
	FullName: "test",
}

var userWithNoPassword = userdomain.User{
	Email:    "test@gmail.com",
	Password: "",
	FullName: "test",
}

var userValid = userdomain.User{
	Email:    "test@gmail.com",
	Password: "testingtesting",
	FullName: "test",
}

var userNotRegistered = userdomain.User{
	Email:    "noexist@gmail.com",
	Password: "testingtesting",
	FullName: "test",
}

var RegisterTestCases = []TestCase{
	{
		Name:       "User with no email(fail)",
		User:       userWithNoEmail,
		StatusCode: 400,
	},
	{
		Name:       "User with no password(fail)",
		User:       userWithNoPassword,
		StatusCode: 400,
	},
	{
		Name:       "User valid(success)",
		User:       userValid,
		StatusCode: 200,
	},
	{
		Name:       "User already exist(fail)",
		User:       userValid,
		StatusCode: 400,
	},
}

var LoginTestCases = []TestCase{
	{
		Name:       "User with no email(fail)",
		User:       userWithNoEmail,
		StatusCode: 400,
	},
	{
		Name:       "User with no password(fail)",
		User:       userWithNoPassword,
		StatusCode: 400,
	},
	{
		Name: "User with invalid password(fail)",
		User: userdomain.User{
			Email:    userValid.Email,
			Password: "password1234",
			FullName: userValid.FullName,
		},
		StatusCode: 401,
	},
	{
		Name:       "User not registered(fail)",
		User:       userNotRegistered,
		StatusCode: 401,
	},
	{
		Name:       "User valid(success)",
		User:       userValid,
		StatusCode: 200,
	},
}

var ResetPasswordTestCases = []TestCase{
	{
		Name:       "User with no email(fail)",
		User:       userWithNoEmail,
		StatusCode: 400,
	},
	{
		Name:       "User with no password(fail)",
		User:       userWithNoPassword,
		StatusCode: 400,
	},
	{
		Name:       "User not registered(fail)",
		User:       userNotRegistered,
		StatusCode: 400,
	},
	{
		Name:       "User valid(success)",
		User:       userValid,
		StatusCode: 200,
	},
}

var ConfirmCodeTestCases = []TestCase{
	{
		Name:       "User with no email(fail)",
		User:       userWithNoEmail,
		StatusCode: 400,
	},
	{
		Name: "User not registered without fullName(fail)",
		User: userdomain.User{
			Email: userNotRegistered.Email,
		},
		StatusCode: 404,
	},
	// code 500 because Email server for sending email
	// is not configured
	{
		Name:       "Valid user(success)",
		User:       userValid,
		StatusCode: 500,
	},
	{
		Name:       "User not registered with fullName(success)",
		User:       userNotRegistered,
		StatusCode: 500,
	},
}

var ContinueWithGoogleTestCases = []TestCase{
	{
		Name:       "User with no email(fail)",
		User:       userWithNoEmail,
		StatusCode: 400,
	},
	{
		Name:       "User valid(success)",
		User:       userValid,
		StatusCode: 200,
	},
	{
		Name:       "User not registered(success)",
		User:       userNotRegistered,
		StatusCode: 200,
	},
}
