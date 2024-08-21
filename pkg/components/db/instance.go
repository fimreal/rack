package db

import (
	"fmt"
	"path"
	"strings"
	"time"

	// "modernc.org/sqlite" // 不好用
	// "gorm.io/driver/sqlite" // 依赖 CGO
	"github.com/fimreal/goutils/ezap"
	"github.com/glebarez/sqlite" // https://github.com/go-gorm/gorm/issues/4101
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 默认配置常量
const (
	DefaultMaxOpenConns    = 300
	DefaultMaxIdleConns    = 30
	DefaultConnMaxLifetime = 1 * time.Hour
	DefaultLogLevel        = logger.Info
)

// DatabaseConfig 表示数据库配置
type DatabaseConfig struct {
	Driver          string          `yaml:"driver"` // 数据库驱动，可选值：mysql, postgres, sqlite
	User            string          `yaml:"user"`
	Password        string          `yaml:"password"`
	Host            string          `yaml:"host"`
	Port            int             `yaml:"port"`
	DBName          string          `yaml:"dbname"`
	MaxOpenConns    int             `yaml:"max_open_conns"`
	MaxIdleConns    int             `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration   `yaml:"conn_max_lifetime"`
	LogLevel        logger.LogLevel `yaml:"log_level"`
}

// NewDatabaseConfig 创建数据库配置并设置默认值
func NewDatabaseConfig(driver, user, password, host string, port int, dbname string) *DatabaseConfig {
	return &DatabaseConfig{
		Driver:          driver,
		User:            user,
		Password:        password,
		Host:            host,
		Port:            port,
		DBName:          dbname,
		MaxOpenConns:    DefaultMaxOpenConns,
		MaxIdleConns:    DefaultMaxIdleConns,
		ConnMaxLifetime: DefaultConnMaxLifetime,
		LogLevel:        DefaultLogLevel, // 使用默认日志级别
	}
}

// NewInstance 创建并返回数据库连接
func NewInstance(config *DatabaseConfig) (*gorm.DB, error) {
	var dsn string
	ezap.Debugf("数据库配置: %+v", config)

	switch config.Driver {
	case "mysql":
		// 构建 MySQL DSN
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
		ezap.Infof("使用 MySQL 数据库: %s", config.DBName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
		}
		return db, nil

	case "postgres":
		// 构建 PostgreSQL DSN
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
		ezap.Infof("使用 PostgreSQL 数据库: %s", config.DBName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
		}
		return db, nil

	case "sqlite":
		// SQLite 不需要用户名和密码, dsn 为数据库文件路径
		if strings.HasPrefix(config.DBName, "/") {
			dsn = config.DBName
		} else {
			workDir := viper.GetString("workdir")
			ezap.Debugf("工作目录: %s", workDir)
			dsn = path.Join(workDir, config.DBName)
		}
		dsn = dsn + ".db"
		ezap.Warnf("使用内 sqlite 数据库: %s", dsn)
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
		}
		return db, nil

	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}
