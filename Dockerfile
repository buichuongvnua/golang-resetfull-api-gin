FROM golang:1.20.4-buster

ENV config=docker
WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -o /bin/app .

CMD env > /app/.env && /bin/app
