FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o ./instacheck .
CMD ./instacheck