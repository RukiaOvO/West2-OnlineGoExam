package dao

import (
	"context"
	"todolist/repository/db/model"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}

	return &UserDao{NewDBClient(ctx)}
}

func (s *UserDao) Create(x *model.UserModel) error {
	return s.db.Create(x).Error
}

func (s *UserDao) FindUserByUserName(userName string) (user *model.UserModel, err error) {
	err = s.db.Model(&model.UserModel{}).
		Where("user_name = ?", userName).
		First(&user).Error
	return
}
func (s *UserDao) FindUserByUserId(userId int64) (user *model.UserModel, err error) {
	err = s.db.Model(&model.UserModel{}).
		Where("id = ?", userId).
		First(&user).Error
	return
}
