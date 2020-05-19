package sqldb

//SQLDBConfg contains all config data necessary to connect MySQL db
type SQLDBConfg struct {
	DriverName string
	UserName   string
	Password   string
	Host       string
	Port       string
	DBName     string
	MaxOpenCon int
	MaxIdleCon int
}
