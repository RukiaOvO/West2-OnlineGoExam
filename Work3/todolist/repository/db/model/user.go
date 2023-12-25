package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"todolist/consts"
)

type UserModel struct {
	ID             int64      `gorm:"column:id; primary_key"`
	UserName       string     `gorm:"column:user_name"`
	PasswordDigest string     `gorm:"column:password_digest"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
	DeletedAt      *time.Time `gorm:"column:deleted_at"`
}

func (*UserModel) TableName() string {
	return "user"
}

func (user *UserModel) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)

	return nil
}

func (user *UserModel) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))

	return err == nil
}
