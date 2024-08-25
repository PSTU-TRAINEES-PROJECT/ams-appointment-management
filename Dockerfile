FROM golang:1.22-alpine

WORKDIR /src

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["air", "serve"]