package gobe

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBDriver string
type DBConnection string

const (
	Mysql    DBDriver     = `mysql`
	Postgres DBDriver     = `postgres`
	Gorm     DBConnection = `gorm`
	Sql      DBConnection = `sql`
)

type GormConnector struct {
	DataSourceName string
	DB             *gorm.DB
}

type SqlConnector struct {
	DataSourceName string
	DB             *sql.DB
}

// Base config is used to initialize connection to an SQL Database
type SqlBaseConfig struct {
	Driver         DBDriver            `mapstructure:"driver" json:"driver"`
	Connector      DBConnection        `mapstructure:"connector" json:"connector"`
	DBName         string              `mapstructure:"db_name" json:"db_name"`
	DBHost         string              `mapstructure:"db_host" json:"db_host"`
	DBPort         string              `mapstructure:"db_port" json:"db_port"`
	DBUsername     string              `mapstructure:"db_username" json:"db_username"`
	DBPassword     string              `mapstructure:"db_password" json:"db_password"`
	SSLMode        string              `mapstructure:"ssl_mode" json:"ssl_mode"`
	DataSourceName string              `mapstructure:"data_source_name" json:"data_source_name"`
	GormConfig     gormConnectorConfig `mapstructure:"gorm" json:"gorm"`
}

// Initialize new connection using pure SQL driver
func NewSqlConfig(config *SqlBaseConfig) SqlConnector {
	dsn := config.DataSourceName
	if dsn == "" {
		switch config.Driver {
		case Mysql:
			dsn = getMySQLConnectionString(config)
		case Postgres:
			dsn = getPostgresConnectionString(config)
		default:
			log.Fatalln("can only initialize DB connection to MySQL or PostgreSQL")
		}
	}

	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	sqlDb := initSqlConnection(config)
	return SqlConnector{DataSourceName: dsn, DB: sqlDb}
}

// Initialize new connection using GORM
func NewGormConfig(config *SqlBaseConfig, table ...interface{}) GormConnector {
	dsn := config.DataSourceName
	if dsn == "" {
		switch config.Driver {
		case Mysql:
			dsn = getMySQLConnectionString(config)
		case Postgres:
			dsn = getPostgresConnectionString(config)
		default:
			log.Fatalln("can only initialize DB connection to MySQL or PostgreSQL")
		}
		config.DataSourceName = dsn
	}

	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	gormDb := initGormConnection(config, table...)
	return GormConnector{DataSourceName: dsn, DB: gormDb}

}

func getMySQLConnectionString(cfg *SqlBaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBName)
}

func getPostgresConnectionString(cfg *SqlBaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.SSLMode)
}

type gormConnectorConfig struct {
	DebugMode       bool `mapstructure:"debug" json:"debug"`
	AutoMigrateMode bool `mapstructure:"auto_migrate" json:"auto_migrate"`
}

func initSqlConnection(baseConfig *SqlBaseConfig) *sql.DB {
	var sqlDb *sql.DB
	switch baseConfig.Driver {
	case Mysql:
		res, err := sql.Open("mysql", baseConfig.DataSourceName)
		if err != nil {
			log.Fatalf("failed to initialize GORM connection to MySQL with error: %s", err.Error())
		}
		sqlDb = res
	case Postgres:
		res, err := sql.Open("postgres", baseConfig.DataSourceName)
		if err != nil {
			log.Fatalf("failed to initialize GORM connection to Postgres with error: %s", err.Error())
		}
		sqlDb = res
	}
	return sqlDb
}

func initGormConnection(baseConfig *SqlBaseConfig, table ...interface{}) *gorm.DB {
	var gormDb *gorm.DB
	switch baseConfig.Driver {
	case Mysql:
		res, err := gorm.Open(mysql.Open(baseConfig.DataSourceName), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to initialize GORM connection to MySQL with error: %s", err.Error())
		}
		gormDb = res
	case Postgres:
		res, err := gorm.Open(postgres.Open(baseConfig.DataSourceName), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to initialize GORM connection to Postgres with error: %s", err.Error())
		}
		gormDb = res
	}
	if baseConfig.GormConfig.DebugMode {
		gormDb.Debug()
	}
	if baseConfig.GormConfig.AutoMigrateMode {
		gormDb.AutoMigrate(table...)
	}
	return gormDb
}
