FROM golang:1.16 AS builder
WORKDIR /go/src/todolist
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /go/src/todolist/cmd/todo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/src/todolist/bin/todo
RUN ls

FROM alpine:3.12.3
RUN adduser -D app-executor
USER app-executor
WORKDIR /app
COPY --from=builder /go/src/todolist/bin/todo /app/todo
COPY --from=builder /go/src/todolist/data/mysql/migrations/todo /app/migrtions

ENV DATABASE_MIGRATIONS_DIR=/app/migrations
EXPOSE 8000

ENTRYPOINT ["/app/todo"]