FROM golang:1.16-alpine

WORKDIR /app

# Add docker-compose-wait tool -------------------
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.2/wait ./wait
RUN chmod +x ./wait

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app cmd/app/main.go

EXPOSE 3000

CMD ["./bin/app"]
