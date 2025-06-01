FROM golang:1.24 AS builder
WORKDIR /app

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - && \
    apt-get install -y nodejs

# Install pnpm
RUN npm install -g pnpm

# Install dependencies
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

COPY . .
RUN go install github.com/a-h/templ/cmd/templ@v0.3.865 &&\
    make build

FROM debian:stable-slim
COPY --from=builder app/bin/goffeine /bin/goffeine
#COPY .env /.env
RUN apt-get update && apt-get install -y ca-certificates
CMD ["/bin/goffeine"]
