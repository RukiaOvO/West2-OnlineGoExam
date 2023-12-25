package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/pkg/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.Success
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(e.InvalidParam, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Token required",
			})
			ctx.Abort()
			return
		}
		claims, err := utils.TokenParse(token)
		if err != nil {
			code = e.Error
			ctx.JSON(e.InvalidParam, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "Failed to parse token",
			})
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(),
			&ctl.UserInfo{Id: claims.Id, UserName: claims.UserName}))
		ctx.Next()
	}
}
