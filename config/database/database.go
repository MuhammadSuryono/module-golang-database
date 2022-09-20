package database

import (
	"gorm.io/gorm"
)

var Connection *gorm.DB

const (
	MysqlConnectionPetra = "mysql_connection_petra"
)

type configOtherDatabase struct {
	config map[string]map[string]interface{}
}

var configMapDatabase = map[string]map[string]interface{}{
	MysqlConnectionPetra: {
		"driver":        InitialDefaultValue("DB_DRIVER_MYSQL", ""),
		"host":          InitialDefaultValue("DB_HOST_MYSQL", ""),
		"password":      InitialDefaultValue("DB_PASS_MYSQL", ""),
		"user":          InitialDefaultValue("DB_USER_MYSQL", ""),
		"database_name": InitialDefaultValue("DB_NAME_MYSQL", ""),
		"port":          InitialDefaultValue("DB_PORT_MYSQL", ""),
		"tables":        MysqlConnectionPetra,
	},
}

type ConnectionHandler struct {
	DB_HOST   string
	DB_PORT   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	TIMEZONE  string
	DB_DRIVER string
}
