package helper

import (
	"app/internal/common/innercall"
	"app/internal/db"
	"app/internal/logic"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func InnerCallLoop() {
	ic := innercall.InnerCallUtil{}.I()
	taskLogic := logic.TaskLogic{}
	for {
		op := ic.Receive()
		logrus.Infof("operationType=%s, content=%s", op.OperationType, string(op.Content))

		switch op.OperationType {
		case innercall.INNERCALL_DELETE:
			{
				taskName := string(op.Content)
				if err := taskLogic.DeleteCronTask(taskName); err != nil {
					logrus.Error(err)
				}
			}

		case innercall.INNERCALL_UPSERT:
			{
				taskData := db.Task{}
				if err := json.Unmarshal(op.Content, &taskData); err != nil {
					logrus.Error(err)
					continue
				}
				if err := taskLogic.UpsertCronTaskToDcron(&taskData); err != nil {
					logrus.Error(err)
				}
			}
		default:
			{
				logrus.Errorf("This operation %s is not support", op.OperationType)
			}
		}
	}
}
