package database

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"os"
)

func (con ConnectionHandler) CreateNewConnection() {
	switch con.DB_DRIVER {
	case "mysql":
		con.mysqlConnection()
		break
	case "postgres":
		con.postgresConnection()
		break
	case "sql-server":
		con.sqlServerConnection()
		break
	default:
		panic("Driver database not set")
	}
}

func CloseConnectionDb(conn *gorm.DB) {
	connDb, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed close database: %v", err))
	}

	_ = connDb.Close()
}

func InitConnectionFromEnvironment() ConnectionHandler {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	return ConnectionHandler{
		DB_HOST:   host,
		DB_PORT:   port,
		DB_USER:   user,
		DB_PASS:   pass,
		DB_NAME:   dbName,
		DB_DRIVER: dbDriver,
	}
}

func InitConnection(host, port, user, pass, dbName, dbDriver string) ConnectionHandler {
	return ConnectionHandler{
		DB_HOST:   host,
		DB_PORT:   port,
		DB_USER:   user,
		DB_PASS:   pass,
		DB_NAME:   dbName,
		DB_DRIVER: dbDriver,
	}
}

func InitOtherConnection() {
	if configMapDatabase == nil {
		panic("Config database another database nil")
	}

	for _, item := range configMapDatabase {
		err := Connection.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open(setDsn(
				os.Getenv(InterfaceToString(item["driver"])),
				os.Getenv(InterfaceToString(item["host"])),
				os.Getenv(InterfaceToString(item["port"])),
				os.Getenv(InterfaceToString(item["user"])),
				os.Getenv(InterfaceToString(item["password"])),
				os.Getenv(InterfaceToString(item["database_name"]))))},
		}, item["tables"]))
		if err != nil {
			panic(fmt.Sprintf("Unable connection to database with error: %v", err))
		}
	}

}

func (configDb configOtherDatabase) CreateOtherConnection() {
	if configDb.config != nil {
		for _, item := range configDb.config {
			err := Connection.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(setDsn(
					InterfaceToString(item["driver"]),
					InterfaceToString(item["host"]),
					InterfaceToString(item["port"]),
					InterfaceToString(item["user"]),
					InterfaceToString(item["password"]),
					InterfaceToString(item["database_name"])))},
			}, item["tables"]))
			if err != nil {
				panic(fmt.Sprintf("Unable connection to database with error: %v", err))
			}
		}
	}
}

func mysqlDsn(host, port, user, pass, dbName string) string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia/Jakarta",
		user, pass, host, port, dbName)
}

func sqlDsn(host, port, user, pass, dbName string) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, pass, host, port, dbName)
}

func pgsqlDsn(host, port, user, pass, dbName string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, pass, dbName, port)
}

func setDsn(driver, host, port, user, pass, dbName string) string {
	fmt.Println(driver, host, port, user, pass, dbName)
	switch driver {
	case "mysql":
		return mysqlDsn(host, port, user, pass, dbName)
	case "postgres":
		return pgsqlDsn(host, port, user, pass, dbName)
	case "sql-server":
		return sqlDsn(host, port, user, pass, dbName)
	default:
		panic(fmt.Sprintf("Driver %v undefined", driver))
	}
	return ""
}

func (con ConnectionHandler) mysqlConnection() {
	args := mysqlDsn(con.DB_HOST, con.DB_PORT, con.DB_USER, con.DB_PASS, con.DB_NAME)
	db, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}

func (con ConnectionHandler) postgresConnection() {
	args := pgsqlDsn(con.DB_HOST, con.DB_PORT, con.DB_USER, con.DB_PASS, con.DB_NAME)
	db, err := gorm.Open(postgres.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}

func (con ConnectionHandler) sqlServerConnection() {
	args := sqlDsn(con.DB_HOST, con.DB_PORT, con.DB_USER, con.DB_PASS, con.DB_NAME)
	db, err := gorm.Open(sqlserver.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}
