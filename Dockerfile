FROM golang:1.22.2

WORKDIR /ascii_art_web_app

COPY go.mod ./
RUN go mod download

COPY . .
CMD ["go", "run", "main.go"]

EXPOSE 8080