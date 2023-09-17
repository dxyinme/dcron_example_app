package innercall

import (
	"app/internal/db"
	"encoding/json"
)

const (
	INNERCALL_DELETE = "DELETE"
	INNERCALL_UPSERT = "UPSERT"
)

type Operation struct {
	OperationType string
	Content       []byte
}

func NewUpsertOperation(task *db.Task) (*Operation, error) {
	b, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}
	return &Operation{
		OperationType: INNERCALL_UPSERT,
		Content:       b,
	}, nil
}

func NewDeleteOperation(taskName string) *Operation {
	return &Operation{
		OperationType: INNERCALL_DELETE,
		Content:       []byte(taskName),
	}
}
