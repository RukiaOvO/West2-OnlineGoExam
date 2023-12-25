package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/pkg/utils"
	"todolist/repository/db/dao"
	"todolist/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(ctx context.Context, req *types.RegisterRequest) (resp interface{}, err error) {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if errors.Is(err, gorm.ErrRecordNotFound) {

		if err = u.SetPassword(req.PassWord); err != nil {
			code = e.ErrorSetPassword
			return ctl.RespError(err, code), err
		}

		u.UserName = req.UserName
		if err = userDao.Create(u); err != nil {
			code = e.ErrorCreateUser
			return ctl.RespError(err, code), err
		}

		return ctl.RespSuccess(e.Success), nil
	} else if err == nil {
		code = e.ErrorExistUser
		err = errors.New("the userId already exists")
		return ctl.RespError(err, code), err
	} else {
		return ctl.RespError(err, e.Error), err
	}
}

func (s *UserSrv) UserLogin(ctx context.Context, req *types.LoginRequest) (resp interface{}, err error) {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	u, err := userDao.FindUserByUserName(req.UserName)
	if err != nil {
		code = e.ErrorNotExistUser
		return ctl.RespError(err, code), err
	}

	if !u.CheckPassword(req.PassWord) {
		code = e.ErrorCheckPassword
		return ctl.RespError(err, code), err
	}

	token, err := utils.TokenGen(u.ID, u.UserName)
	if err != nil {
		code = e.ErrorTokenGen
		return ctl.RespError(err, code), err
	}

	userResp := types.TokenData{
		User: types.UserLoginResponse{
			Id:        u.ID,
			UserName:  u.UserName,
			CreatedAt: u.CreatedAt.Unix(),
		},
		AccessToken: token,
	}

	return ctl.RespSuccessWithData(userResp, code), nil
}
