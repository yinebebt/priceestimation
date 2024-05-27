FROM golang:latest

WORKDIR /app


COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

# Build the app
RUN go build -o app cmd/main.go

EXPOSE 8082

CMD ["/app/wait-for-it.sh", "db:5432", "--","go","run","./cmd/main.go"]