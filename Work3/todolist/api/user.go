package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/pkg/utils"
	"todolist/service"
	"todolist/types"
)

func RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.RegisterRequest
		if err := ctx.ShouldBind(&req); err == nil {
			tmp := service.GetUserSrv()
			resp, regErr := tmp.UserRegister(ctx.Request.Context(), &req)
			if regErr != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(regErr))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.LoginRequest
		if err := ctx.ShouldBind(&req); err == nil {
			tmp := service.GetUserSrv()
			resp, err := tmp.UserLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
