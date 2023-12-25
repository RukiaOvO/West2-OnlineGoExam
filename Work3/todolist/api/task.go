package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todolist/consts"
	"todolist/pkg/utils"
	"todolist/service"
	"todolist/types"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskRequest
		if err := ctx.ShouldBind(&req); err == nil {
			tmp := service.GetTaskSrv()
			resp, err := tmp.CreateTask(ctx.Request.Context(), &req)
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

func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ListTaskRequest
		if err := ctx.ShouldBind(&req); err == nil {
			if req.Limit <= 0 {
				req.Limit = consts.DefaultTaskLimit
			}
			tmp := service.GetTaskSrv()
			resp, err := tmp.ListTask(ctx.Request.Context(), &req)
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

func ShowTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		tmp := service.GetTaskSrv()
		resp, err := tmp.ShowTask(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

func UpdateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UpdateTaskRequest
		if err := ctx.ShouldBind(&req); err == nil {
			id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
			tmp := service.GetTaskSrv()
			resp, err := tmp.UpdateTask(ctx.Request.Context(), &req, id)
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

func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskRequest
		if err := ctx.ShouldBind(&req); err == nil {
			if req.Limit <= 0 {
				req.Limit = consts.DefaultTaskLimit
			}
			tmp := service.GetTaskSrv()
			resp, err := tmp.SearchTask(ctx.Request.Context(), &req)
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

func DeleteTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			utils.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
		tmp := service.GetTaskSrv()
		resp, err := tmp.DeleteTask(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
