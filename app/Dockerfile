FROM golang:1.19.5-alpine

WORKDIR /app
COPY ./app/go.sum ./
COPY ./app/go.mod ./

RUN go mod download
RUN go mod verify

COPY ./app .
RUN go build -o main main.go

CMD ["./main"]
