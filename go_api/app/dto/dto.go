package dto

type LightDto struct {
	Name   string `form:"name" binding:"required"`
	Switch *int   `form:"switch" binding:"required"`
}

type Operate struct {
	ShowId    int    `form:"showId" binding:"required"`
	Id        int    `form:"id" binding:"required"`
	Ip        string `form:"ip" binding:"required"`
	Operation string `form:"operation" binding:"required"`
	Offset    int    `form:"offset"`
}

type Status struct {
	Ip     string `json:"ip" binding:"required"`
	ShowId int    `json:"showId" binding:"required"`
}

type IsPlay struct {
	Data   []Status `json:"data" binding:"required"`
	Status *int     `json:"status" binding:"required"`
}

type DvdDto struct {
	Status string `form:"status" binding:"required"`
}

type Login struct {
	ErrInfo string `json:"errInfo"`
	Project string `json:"project"`
	Rescode string `json:"rescode"`
	Token   string `json:"token"`
}
