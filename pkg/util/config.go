package util

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port      string
	SecretKey string
	Database  Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() Config {
	return Config{
		Port:      os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		Database:  LoadDatabase(),
	}
}

func LoadDatabase() Database {
	return Database{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_DB"),
	}
}

func SetupDB(config Config) *gorm.DB {
	dsn :=
		"host=" + config.Database.Host +
			" user=" + config.Database.User +
			" password=" + config.Database.Password +
			" dbname=" + config.Database.Name +
			" port=" + config.Database.Port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}
