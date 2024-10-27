package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка при загрузке файла конфигурации: %v", err)
	}
}

func initDB() {
	initConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
	)

	fmt.Println(connStr)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS tickers (
		id SERIAL PRIMARY KEY,
		symbol VARCHAR(50),
		price REAL,
		timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}
