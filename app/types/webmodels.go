package types

type TaskBase struct {
	CronStr string `json:"cronStr" binding:"required"`
	SQLStr  string `json:"SQLStr" binding:"required"`
	DBName  string `json:"dbName" binding:"required"`
}

type DBBase struct {
	DBType       string `json:"dbType" binding:"required"`
	User         string `json:"user" binding:"required"`
	Password     string `json:"password" binding:"required"`
	DatabaseName string `json:"databaseName" binding:"required"`
	Addr         string `json:"addr" binding:"required"`
}

// Task example
type Task struct {
	Name string `json:"name"`
	TaskBase
}

// TaskReq example
type TaskReq struct {
	TaskBase
}

// RunSQLReq example
type RunSQLReq struct {
	DBCustomName string `json:"dbCustomName" binding:"required"`
	SQLStr       string `json:"SQLStr" binding:"required"`
}

// DBReq example
type DBReq struct {
	DBBase
}

// DB example
type DB struct {
	CustomName string `json:"customName" binding:"required"`
	DBBase
}
