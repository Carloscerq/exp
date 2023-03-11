FROM golang:1.20.1-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY . .

RUN go build -o ./app 

EXPOSE 3000

CMD  ./app -action up
