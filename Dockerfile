FROM golang:1.24.9-alpine

WORKDIR /app

RUN apk add --no-cache git

RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH=$PATH:/go/bin

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY src/ .
RUN swag init -g main.go --parseDependency --output docs

RUN go build -o main

EXPOSE 8080

CMD ["./main"]