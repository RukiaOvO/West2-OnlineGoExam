package ctl

import (
	"net/http"
	"todolist/pkg/e"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

func RespList(items interface{}, total int64) Response {
	return Response{
		Status: e.Success,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "操作成功",
	}
}
func RespSuccess(code int) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

func RespSuccessWithData(data interface{}, code int) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   data,
		Msg:    e.GetMsg(code),
		Error:  "",
	}
}

func RespError(err error, code int) *Response {
	return &Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}
