FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]