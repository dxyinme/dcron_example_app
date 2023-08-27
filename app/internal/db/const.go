package db

type DBType = uint

const (
	DBTypeInvalid    = 0
	DBTypeMYSQL      = 1
	DBTypePOSTGRESQL = 2

	MySQLDSNFormat      = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	PostgreSQLDSNFormat = ""
)

var (
	DBTypeMap = map[string]DBType{
		"mysql":      DBTypeMYSQL,
		"postgresql": DBTypePOSTGRESQL,
	}

	ReserveDBTypeMap = map[DBType]string{
		DBTypeMYSQL:      "mysql",
		DBTypePOSTGRESQL: "postgresql",
	}
)
