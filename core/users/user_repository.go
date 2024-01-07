package users

type UserRepository interface {
	Save(user User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
}
