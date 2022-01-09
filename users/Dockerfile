FROM golang:alpine as builder
ARG EXPOSED_PORT=$USER_SERVICE_PORT
EXPOSE ${EXPOSED_PORT}
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./user-service

FROM alpine
WORKDIR app
COPY --from=builder /app/user-service ./user-service
COPY --from=builder /app/models/migrations ./migrations

CMD ["/app/user-service"]