package userdomain

type Repository interface {
	GetByEmail(email string) (*User, error)
	Register(user *User) error
	ResetPassword(user *User) error
}
