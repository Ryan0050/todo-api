package config

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type EnvGormConfig struct {
	LogLevel int `env:"GORM_LOG_LEVEL" envDefault:"1"`
}

type EnvPostgreSQLConfig struct {
	PostgresSQLAppName    string `env:"POSTGRE_SQL_APP_NAME"`
	PostgreSQLHost        string `env:"POSTGRE_SQL_HOST"`
	PostgreSQLPort        int    `env:"POSTGRE_SQL_PORT"`
	PostgreSQLUser        string `env:"POSTGRE_SQL_USER"`
	PostgreSQLPassword    string `env:"POSTGRE_SQL_PASSWORD"`
	PostgreSQLDBName      string `env:"POSTGRE_SQL_DB_NAME"`
	PostgreSQLDBSchema    string `env:"POSTGRE_SQL_DB_SCHEMA" envDefault:"public"`
	PostgreSQLDBLogName   string `env:"POSTGRE_SQL_DB_LOG_NAME"`
	PostgreSQLDBLogSchema string `env:"POSTGRE_SQL_DB_LOG_SCHEMA"`
}

var (
	PostgreSQLConfig EnvPostgreSQLConfig
	GormConfig       EnvGormConfig
)

type connectionMapper struct {
	TenantConnection map[string][]*gorm.DB
	sync.Mutex
}

var ConnectionMapper *connectionMapper
var connectionOnce sync.Once

func InitConnectDB(dbHost, dbUser, dbPass, dbName, prefix string, dbPort int) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta application_name=%s", 
		dbHost, dbUser, dbPass, dbName, dbPort, PostgreSQLConfig.PostgresSQLAppName)
	
	log.Println("Connecting with DSN:", dsn)
	
	var tablePrefix string
	if prefix != "" {
		tablePrefix = prefix + "."
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(GormConfig.LogLevel)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
			TablePrefix:   tablePrefix,
		},
	})

	if err != nil {
		log.Println("Database connection error:", err)
		return nil
	}

	log.Println("Connected to Postgres DB")
	return db
}

func GetConnectionMapper() *connectionMapper {
	connectionOnce.Do(func() {
		ConnectionMapper = &connectionMapper{
			TenantConnection: make(map[string][]*gorm.DB),
			Mutex:            sync.Mutex{},
		}
	})

	return ConnectionMapper
}