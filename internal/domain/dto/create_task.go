package dto

type CreateTaskReq struct {
	Title       string      `json:"title" validate:"required"`
	Description string      `json:"description" validate:"required"`
	Deadline    DatetimeReq `json:"deadline"`
}

type CreateTaskResp struct {
	ID int `json:"id"`
}
