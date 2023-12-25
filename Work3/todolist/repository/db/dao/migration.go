package dao

import (
	"fmt"
	"todolist/repository/db/model"
)

func structMigrate() {
	if err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.UserModel{}, &model.TaskModel{}); err != nil {
		panic(err)
	}

	fmt.Println("Migrated successfully")
}
