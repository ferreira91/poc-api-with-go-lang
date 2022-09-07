package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DriverName string
	Host       string
	Port       int
	User       string
	Password   string
	DbName     string
	SslMode    string
}

func Init() Config {
	return Config{
		DriverName: viper.GetString("DB_DRIVER_NAME"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetInt("DB_PORT"),
		User:       viper.GetString("DB_USER"),
		Password:   viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
		SslMode:    viper.GetString("DB_SSL_MODE"),
	}
}

func (c *Config) Start() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DbName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")
	return db, nil
}
