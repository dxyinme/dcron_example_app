package crontasks

type CronTask struct {
	Name           string
	CronStr        string
	DBCustomerName string
	SQLStr         string
}
