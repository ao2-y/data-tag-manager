package mysql

import (
	"ao2-y/data-tag-manager/logger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

func NewDBConnection(
	Host string,
	Port int,
	User string,
	Password string,
	DatabaseName string,
) *gorm.DB {
	return getMysqlConn(Host, Port, User, Password, DatabaseName)
}

type gormLog struct {
	zap *zap.Logger
}

func (l *gormLog) Printf(msg string, values ...interface{}) {
	l.zap.Info(fmt.Sprintf(msg, values...))
}

func getMysqlConn(
	Host string,
	Port int,
	User string,
	Password string,
	DatabaseName string,
) *gorm.DB {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?parseTime=true&loc=Local",
		User,
		Password,
		Host,
		Port,
		DatabaseName,
	)
	// TODO 暫定
	logger := logger.InitApplicationLogger()
	gormLog := &gormLog{zap: logger}
	newLogger := gormLogger.New(
		gormLog,
		gormLogger.Config{
			SlowThreshold: 5 * time.Second, // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      false,           // Disable color
		},
	)
	conn, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("database connection error.:%w", err))
	}
	return conn
}
