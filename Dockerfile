FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o goffeine .

FROM debian:stable-slim
COPY --from=builder /app/goffeine /bin/goffeine
COPY .env /.env
RUN apt-get update && apt-get install -y ca-certificates
CMD ["/bin/goffeine"]