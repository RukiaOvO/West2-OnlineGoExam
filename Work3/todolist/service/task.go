package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
	"todolist/pkg/ctl"
	"todolist/pkg/e"
	"todolist/repository/db/dao"
	"todolist/repository/db/model"
	"todolist/types"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskRequest) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	if err != nil {
		return ctl.RespError(err, e.ErrorNotExistUser), err
	}
	task := &model.TaskModel{
		Uid:       user.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    req.Status,
		StartTime: time.Now().Unix(),
	}
	if err = dao.NewTaskDao(ctx).CreateTask(task); err != nil {
		return ctl.RespError(err, e.ErrorCreateTask), err
	}

	return ctl.RespSuccess(e.Success), nil
}

func (s *TaskSrv) ListTask(ctx context.Context, req *types.ListTaskRequest) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, u.Id)
	if err != nil {
		return ctl.RespError(err, e.ErrorDatabase), err
	}
	tRespList := make([]*types.ListTaskResponse, 0)
	for _, v := range tasks {
		tRespList = append(tRespList, &types.ListTaskResponse{
			Id:        v.Id,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	return ctl.RespList(tRespList, total), nil
}

func (s *TaskSrv) ShowTask(ctx context.Context, id int64) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserId(id, u.Id)
	if err != nil {
		return ctl.RespError(err, e.ErrorFindTask), err
	}
	task.AddView() //展示后+1
	tResp := &types.ListTaskResponse{
		Id:        task.Id,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(),
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}

	return ctl.RespSuccessWithData(tResp, e.Success), nil
}

func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskRequest, id int64) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}

	if err = dao.NewTaskDao(ctx).UpdateTask(u.Id, req, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return ctl.RespError(err, e.ErrorUpdateTask), err
	}

	return ctl.RespSuccess(e.Success), nil
}

func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskRequest) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}
	tasks, total, err := dao.NewTaskDao(ctx).SearchTask(u.Id, req)
	if err != nil {
		return ctl.RespError(err, e.ErrorSearchTask), err
	}
	taskList := make([]*types.ListTaskResponse, 0)
	for _, v := range tasks {
		taskList = append(taskList, &types.ListTaskResponse{
			Id:        v.Id,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	return ctl.RespList(taskList, total), nil
}

func (s *TaskSrv) DeleteTask(ctx context.Context, id int64) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		return ctl.RespError(err, e.ErrorGetUserInfo), err
	}

	if err = dao.NewTaskDao(ctx).DeleteTask(u.Id, id); errors.Is(err, gorm.ErrRecordNotFound) {
		return ctl.RespError(err, e.ErrorDeleteTask), err
	}

	return ctl.RespSuccess(e.Success), nil
}
