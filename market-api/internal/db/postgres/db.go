package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Config struct {
	DriverName string
	Host       string
	Port       string
	User       string
	Password   string
	DbName     string
	SslMode    string
}

func Init() Config {
	return Config{
		DriverName: viper.GetString("DB_DRIVER_NAME"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetString("DB_PORT"),
		User:       viper.GetString("DB_USER"),
		Password:   viper.GetString("DB_PASSWORD"),
		DbName:     viper.GetString("DB_NAME"),
		SslMode:    viper.GetString("DB_SSL_MODE"),
	}
}

func (c *Config) Start() *sql.DB {
	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DbName, c.SslMode)

	db, err := sql.Open(c.DriverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
