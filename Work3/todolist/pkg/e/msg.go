package e

var MsgFlags = map[int]string{
	Success:            "操作成功", //common errors
	Error:              "操作失败",
	InvalidParam:       "请求参数错误",
	ErrorGetUserInfo:   "获取用户信息错误",
	ErrorUnmarshalJson: "解码Json错误",
	ErrorDatabase:      "数据库错误",

	ErrorCreateTask: "创建事项时错误", //task errors
	ErrorFindTask:   "查找事项时错误",
	ErrorUpdateTask: "更新事项时错误",
	ErrorDeleteTask: "删除事项时错误",
	ErrorSearchTask: "匹配事项时错误",
	ErrorListTask:   "列出事项时错误",

	ErrorExistUser:     "用户存在", //user errors
	ErrorNotExistUser:  "用户不存在",
	ErrorCreateUser:    "创建用户时错误",
	ErrorSetPassword:   "设置密码时错误",
	ErrorCheckPassword: "检查密码时错误",
	ErrorTokenGen:      "Token生成错误",
	ErrorCheckToken:    "Token检查错误",
	ErrorParseToken:    "Token解析错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}

	return msg
}
