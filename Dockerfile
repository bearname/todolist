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
##
#FROM golang:1.16 AS modules
#
#COPY ./go.mod ./go.sum /
#RUN go mod download
#
#FROM golang:1.16 AS builder
#RUN useradd -u 1001 appuser
#
#COPY --from=modules /go/pkg go/pkg
#COPY . /build
#WORKDIR /build
#
#RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 \
#    go build -o ./bin/todo ./cmd/todo
#
#RUN chmox +x ./bin/todo
#
#FROM scratch
#
#COPY --from=builder /etc/passwd /etc/passwd
#USER appuser
#
#COPY --from=builder /build/bin/todo /app/bin/todo
#COPY --from=builder /build/data/mysql/migrations/todo /app/migrtions
#ENV DATABASE_MIGRATIONS_DIR=/app/migrations
#
#EXPOSE 8080
#
#CMD ["/app/bin/budget"]
#
# Download modules
#FROM golang:1.16 AS modules
#
#COPY ./go.mod ./go.sum /
#RUN go mod download
#
#
## Build the binary
#FROM golang:1.16 AS builder
#
#RUN useradd -u 1001 appuser
#
#COPY --from=modules /go/pkg /go/pkg
#COPY . /build
#WORKDIR /build
#
#RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 \
#    go build -o ./bin/todo ./cmd/todo
#
#RUN chmod +x ./bin/todo
#
#
## Run the binary
#FROM scratch
#
#COPY --from=builder /etc/passwd /etc/passwd
#USER appuser
#
#COPY --from=builder /build/bin/todo /app/bin/todo
#
#COPY --from=builder /build/data/mysql/migrations/todo /app/migrations
#ENV DATABASE_MIGRATIONS_DIR=/app/migrations
#
#EXPOSE 8000
#
#CMD ["/app/bin/todo"]