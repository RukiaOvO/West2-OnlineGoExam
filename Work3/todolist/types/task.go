package types

type CreateTaskRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int64  `json:"status" form:"status"`
}

type UpdateTaskRequest struct {
	NewStatus int64 `json:"new_status" form:"new_status"`
}

type SearchTaskRequest struct {
	KeyWord string `json:"keyword" form:"keyword"`
	Limit   int    `json:"limit" form:"limit"`
	Start   int    `json:"start" form:"start"`
}

type ListTaskRequest struct {
	Limit int `json:"limit" form:"limit"`
	Start int `json:"start" form:"start"`
}

type ListTaskResponse struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content" `
	View      int64  `json:"view"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
