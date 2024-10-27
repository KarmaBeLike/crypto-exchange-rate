package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/rate", getCurrentRate)
	r.GET("/history", getRateHistory)

	return r
}

func getCurrentRate(c *gin.Context) {
	symbol := c.Query("symbol")

	pageSize, offset := getPaginationParams(c)

	var rows *sql.Rows
	var err error

	if symbol != "" {
		rows, err = db.Query(
			`SELECT id, symbol, price, timestamp 
             FROM tickers 
             WHERE symbol = $1 
             ORDER BY timestamp DESC 
             LIMIT 1`, symbol)
	} else {
		rows, err = db.Query(
			`SELECT DISTINCT ON (symbol) id, symbol, price, timestamp 
             FROM tickers 
             ORDER BY symbol, timestamp DESC
             LIMIT $1 OFFSET $2`, pageSize, offset)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var tickers []Ticker
	for rows.Next() {
		var ticker Ticker
		err := rows.Scan(&ticker.ID, &ticker.Symbol, &ticker.Price, &ticker.Timestamp)
		if err != nil {
			continue
		}
		tickers = append(tickers, ticker)
	}

	c.JSON(http.StatusOK, gin.H{
		"page_size":   pageSize,
		"total_items": len(tickers),
		"tickers":     tickers,
	})
}

func getRateHistory(c *gin.Context) {
	symbol := c.Query("symbol")
	limit := time.Now().Add(-24 * time.Hour)

	pageSize, offset := getPaginationParams(c)

	rows, err := db.Query(
		`SELECT id, symbol, price, timestamp 
	FROM tickers 
	WHERE symbol = $1 AND timestamp >= $2
	ORDER BY timestamp DESC  
	LIMIT $3 OFFSET $4`, symbol, limit, pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var tickers []Ticker
	for rows.Next() {
		var ticker Ticker
		err := rows.Scan(&ticker.ID, &ticker.Symbol, &ticker.Price, &ticker.Timestamp)
		if err != nil {
			continue
		}
		tickers = append(tickers, ticker)
	}

	c.JSON(http.StatusOK, gin.H{
		"page_size":   pageSize,
		"total_items": len(tickers),
		"tickers":     tickers,
	})
}
