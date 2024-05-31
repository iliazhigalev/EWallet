FROM golang:1.20.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000