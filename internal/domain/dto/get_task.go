package dto

import "go-service-template/pkgs/utils/pagable"

type GetTaskByIdReq struct {
	ID int `params:"id" validate:"required"`
}

type GetTaskByIdResp struct {
	TaskDTO
}

type GetTaskListReq struct {
	Query *pagable.Query
}

type GetTaskListResp struct {
	pagable.ListResponse
}
