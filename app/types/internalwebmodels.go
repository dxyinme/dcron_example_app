package types

type RemoveTaskReq struct {
	Name string `json:"name" binding:"required"`
}
