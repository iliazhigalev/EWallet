FROM golang:1.22.3 AS builder

WORKDIR /usr/src/app

# Копируем зависимости и собираем бинарный файл
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server cmd/server/main.go

# Этап 2: Запуск приложения
FROM alpine:latest
COPY --from=builder /usr/src/app/server .
COPY .env .


# Открываем порт для приложения
EXPOSE 3000

# Запускаем сервер
CMD ["./server"]
