FROM golang:alpine as builder
ARG EXPOSED_PORT=$USER_SERVICE_PORT
EXPOSE ${EXPOSED_PORT}
WORKDIR /app
COPY . ./

RUN go mod download

COPY *.go ./
RUN go build -o ./user-service

FROM alpine
COPY --from=builder /app/user-service .
CMD ["/user-service"]