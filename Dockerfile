FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN make build
