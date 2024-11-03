package domain

type User struct {
	UserID   string `json:"user_id" gorm:"primary_key"`
	Username string `json:"name" gorm:"varchar(255);unique"`
	Role     string `json:"role" gorm:"varchar(50)"`
	Password string `json:"password" gorm:"varchar(255)"`
}
