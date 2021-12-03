FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o rest-server cmd/rest-server/main.go 

EXPOSE 1234

CMD ["./rest-server"]