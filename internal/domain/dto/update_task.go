package dto

type UpdateTaskReq struct {
	ID          int         `params:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Deadline    DatetimeReq `json:"deadline"`
}

type UpdateTaskStatusReq struct {
	ID     int `params:"id"`
	Status int `json:"status"`
}
