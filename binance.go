package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"crypto-exchange-rate/values"

	"github.com/gorilla/websocket"
)

func connectToBinance() {
	client, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/!ticker@arr", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket:", err)
	}
	defer client.Close()

	for {
		_, message, err := client.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		var tickers []map[string]interface{}
		if err := json.Unmarshal(message, &tickers); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		saveTickers(tickers)
	}
}

func saveTickers(tickers []map[string]interface{}) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error beginning transaction:", err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO tickers(symbol, price) VALUES($1, $2)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		tx.Rollback()
		return
	}
	defer stmt.Close()

	for _, t := range tickers {
		symbol := t[values.Symbol].(string)
		if !strings.HasSuffix(symbol, "USDT") {
			continue
		}
		trimmedSymbol := strings.TrimSuffix(symbol, "USDT")
		priceStr := t[values.Price].(string)
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error parsing price:", err)
			continue
		}

		_, err = stmt.Exec(trimmedSymbol, price)
		if err != nil {
			log.Println("Error inserting ticker:", err)
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction:", err)
		tx.Rollback()
		return
	}
}
