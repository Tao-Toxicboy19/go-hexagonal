package users

type User struct {
	ID    uint   `gorm:"AUTO_INCREMENT" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}