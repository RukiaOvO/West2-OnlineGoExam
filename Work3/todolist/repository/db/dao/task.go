package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"todolist/repository/db/model"
	"todolist/types"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}

	return &TaskDao{NewDBClient(ctx)}
}

func (s *TaskDao) ListTask(start int, limit int, uId int64) (r []*model.TaskModel, total int64, err error) {
	if err = s.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId).
		Count(&total).Error; err != nil {
		return
	}

	err = s.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId).
		Limit(limit).
		Offset((start - 1) * limit).
		Find(&r).Error
	return
}

func (s *TaskDao) FindTaskByIdAndUserId(id, userId int64) (r *model.TaskModel, err error) {
	err = s.db.Model(&model.TaskModel{}).
		Where("id = ? AND uid = ?", id, userId).
		First(&r).Error
	return
}

func (s *TaskDao) CreateTask(x *model.TaskModel) error {
	return s.db.Create(x).Error
}

func (s *TaskDao) UpdateTask(uId int64, req *types.UpdateTaskRequest, id int64) (err error) {
	tTask := &model.TaskModel{}
	dao := s.db.Model(&model.TaskModel{}).
		Where("uid = ? ", uId)
	if err := dao.
		Where("id = ?", id).
		First(&tTask).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	tTask.Status = req.NewStatus
	err = s.db.Save(tTask).Error
	return
}

func (s *TaskDao) SearchTask(uId int64, req *types.SearchTaskRequest) (r []*model.TaskModel, total int64, err error) {
	dao := s.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId) //uid预定位，限制只显示自己的task

	if err = dao.
		Where("title LIKE ? OR content LIKE ?", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%").
		Count(&total).Error; err != nil {
		return
	}

	err = dao.
		Where("title LIKE ? OR content LIKE ?", "%"+req.KeyWord+"%", "%"+req.KeyWord+"%").
		Limit(req.Limit).
		Offset((req.Start - 1) * req.Limit).
		Find(&r).Error
	return
}

func (s *TaskDao) DeleteTask(uId int64, id int64) (err error) {
	var tmp *model.TaskModel

	dao := s.db.Model(&model.TaskModel{}).
		Where("uid = ?", uId) //uid预定位，限制只删除自己的task
	if err := dao.Where("id = ?", id).First(&tmp).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = dao.
		Where("id = ?", id).
		Delete(&tmp).Error
	return
}
