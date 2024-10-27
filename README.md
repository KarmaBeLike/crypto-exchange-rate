# Crypto Exchange Rate Service

## Introduction
This project is a Go application that fetches real-time cryptocurrency exchange rates from Binance and saves them to a PostgreSQL database to see a history. The project is containerized using Docker and configured to run with Docker Compose for easy setup and management.


## Technologies

- **Go** 1.23
- **PostgreSQL** (latest version)
- **Docker** and **Docker Compose**

##  Installation 
1. Clone the repository:
   ```bash
   git clone https://github.com/KarmaBeLike/crypto-exchange-rate.git


### START PROJECT
- **install dependencies:**
```
go mod tidy
```
- **run project:**
```
docker-compose up --build
```
- **to manually enter the database:**
```
docker exec -it crypto_postgres psql -U postgres -d postgres

```
- **Fetching crypto exchange rate of all tickers with a pagination:**
    ```http
    GET /rate
    ```
    - sample output:
    ```json
    {
    "page_size": 10,
    "tickers": [
        {
            "id": 166061,
            "symbol": "1000SATS",
            "price": 0.000238,
            "timestamp": "2024-10-27T13:10:05.41959Z"
        },
        {
            "id": 165988,
            "symbol": "1INCH",
            "price": 0.2535,
            "timestamp": "2024-10-27T13:10:05.41959Z"
        },
        {
            "id": 165927,
            "symbol": "1MBABYDOGE",
            "price": 0.0025757,
            "timestamp": "2024-10-27T13:10:04.418531Z"
        },
        {
            "id": 165981,
            "symbol": "AAVE",
            "price": 143.14,
            "timestamp": "2024-10-27T13:10:05.41959Z"
        }
        ...
    }
    ```
     ```http
    GET /rate?page=5
    ```
    - sample output:
    ```json
    {
    "page_size": 10,
    "tickers": [
        {
            "id": 186524,
            "symbol": "ATOM",
            "price": 4.366,
            "timestamp": "2024-10-27T13:11:44.445157Z"
        },
        {
            "id": 186599,
            "symbol": "AUCTION",
            "price": 13.14,
            "timestamp": "2024-10-27T13:11:44.445157Z"
        },
        {
            "id": 183763,
            "symbol": "AUDIO",
            "price": 0.1203,
            "timestamp": "2024-10-27T13:11:30.447607Z"
        }
        ...
    }
    ```
- **Fetching crypto exchange rate of a specific ticker:**
    ```http
    GET /rate?symbol=BTC
    ```
    - sample output:
    ```json
   {
    "page_size": 10,
    "tickers": [
        {
            "id": 256776,
            "symbol": "BTC",
            "price": 67620.94,
            "timestamp": "2024-10-27T13:17:44.590856Z"
        }
    ],
    "total_items": 1
}

- **Get the history of changes for a specific ticker with a pagination:**
    ```http
    GET /history?symbol=BTC&page=5
    ```
    - sample output:
    ```json
   {
    "page_size": 10,
    "tickers": [
        {
            "id": 318215,
            "symbol": "BTC",
            "price": 67699.99,
            "timestamp": "2024-10-27T13:22:51.734756Z"
        },
        {
            "id": 317997,
            "symbol": "BTC",
            "price": 67699.99,
            "timestamp": "2024-10-27T13:22:50.721837Z"
        },
        {
            "id": 317760,
            "symbol": "BTC",
            "price": 67699.99,
            "timestamp": "2024-10-27T13:22:49.737596Z"
        },
        {
            "id": 317453,
            "symbol": "BTC",
            "price": 67700,
            "timestamp": "2024-10-27T13:22:48.735534Z"
        }
        ...
   }
---