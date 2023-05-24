FROM golang:1.19.3

WORKDIR /app

COPY . ./

RUN go build -o main cmd/app/main.go

EXPOSE 8080:8080

CMD ["/app/main"]