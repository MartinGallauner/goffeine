FROM debian:stable-slim
COPY goffeine /bin/goffeine
COPY .env /.env
RUN apt-get update && apt-get install -y ca-certificates
CMD ["/bin/goffeine"]