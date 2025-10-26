FROM golang:1.25-alpine

WORKDIR /src/app

# Install Air for live reload
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy
