package model

import (
	"github.com/spf13/cast"
	"time"
	"todolist/repository/cache"
)

type TaskModel struct {
	Id        int64      `gorm:"column:id; primary_key"`
	Uid       int64      `gorm:"column:uid"`
	Title     string     `gorm:"column:title"`
	Content   string     `gorm:"column:content"`
	Status    int64      `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
	StartTime int64      `gorm:"column:start_time"`
	EndTime   int64      `gorm:"column:end_time"`
}

func (*TaskModel) TableName() string {
	return "task"
}

func (t *TaskModel) View() int64 {
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(t.Id)).Result()
	return cast.ToInt64(countStr)
}

func (t *TaskModel) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(t.Id))
}
