package models

import (
	"time"
	"tourism/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	ID           uint       `gorm:"primary_key" josn:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"deleted_at"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	NationalCode string     `gorm:"index;type:nvarchar(10)" json:"national_code"`
	IsActive     *bool      `gorm:"default:true" json:"is_active"`
	Mobile       string     `gorm:"unique;not null;type:nvarchar(15)" json:"mobile"`
	Password     string     `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(userId int) User {
	var user User
	db.First(&user, userId)
	return user
}

func CreateUser(user *User) *gorm.DB {
	result := db.Create(&user)

	return result
}

func UpdateUser(user *User, userId int) *gorm.DB {
	newData := User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		NationalCode: user.NationalCode,
		IsActive:     user.IsActive,
		Mobile:       user.Mobile,
		Password:     user.Password, // TODO Hash.
	}

	result := db.Model(User{}).Where("ID = ?", userId).Updates(&newData)

	return result
}

func DeleteUser(userId int) {
	db.Delete(&User{}, userId)
}
