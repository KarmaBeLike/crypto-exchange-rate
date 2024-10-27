# Используем официальный образ Go
FROM golang:1.23

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Собираем приложение
RUN go build -o main .

# Открываем порт для вашего приложения
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
