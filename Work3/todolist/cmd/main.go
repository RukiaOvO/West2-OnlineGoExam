package main

import (
	"todolist/conf"
	"todolist/pkg/utils"
	"todolist/repository/cache"
	"todolist/repository/db/dao"
	"todolist/routes"
)

func main() {
	load()
	r := routes.NewRouter()
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}

func load() {
	conf.InitConfig()
	dao.InitDb()
	utils.InitLog()
	cache.RedisInit()
}
