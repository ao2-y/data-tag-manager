package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(
	Host string,
	Port string,
	User string,
	Password string,
	DatabaseName string,
) *gorm.DB {
	return getMysqlConn(Host, Port, User, Password, DatabaseName)
}

func getMysqlConn(
	Host string,
	Port string,
	User string,
	Password string,
	DatabaseName string,
) *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		User,
		Password,
		Host,
		Port,
		DatabaseName,
	)
	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("database connection error.:%w", err))
	}
	return conn
}
