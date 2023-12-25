package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todolist/api"
	"todolist/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	reqs := r.Group("/api")
	{
		reqs.GET("test", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		reqs.POST("user/register", api.RegisterHandler())
		reqs.POST("user/login", api.LoginHandler())

		authed := reqs.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTaskHandler())       //增
			authed.GET("tasks", api.ListTaskHandler())         //列表
			authed.GET("task/:id", api.ShowTaskHandler())      //查
			authed.PUT("task/:id", api.UpdateTaskHandler())    //改
			authed.POST("search", api.SearchTaskHandler())     //模糊匹配
			authed.DELETE("task/:id", api.DeleteTaskHandler()) //删
		}
	}
	return r
}
