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
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, 5432, "postgres", "postgres", "market")
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	log.Println("Database connection established")
	return db

	//connStr := fmt.Sprintf(
	//	"host=%s port=%d user=%s password=%s  dbname=%s sslmode=disable",
	//	"db", 5431, "postgres", "postgres", "market",
	//)
	//db2, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatal(err.Error())
	//	print(db2.Ping())
	//}
	//println(db2.Ping())
	//return db2

	//var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	c.Host, c.Port, c.User, c.Password, c.DbName, c.SslMode)
	//println(dataSourceName)
	//log.Printf("String connection: %s", dataSourceName)
	//
	//db, err := sql.Open("postgers", dataSourceName)
	//if err != nil {
	//	println(err)
	//	log.Fatal(err)
	//	return nil
	//}
	//err = db.Ping()
	//if err != nil {
	//	println(err)
	//	log.Fatal(err)
	//	return nil
	//}
	//log.Println("Database connection established")
	//return db
}
