# Etapa 1: build
FROM golang:1.24.2 AS builder

LABEL maintainer="Ricardo <florez.ricardo.2748@eam.edu.co>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

# Etapa 2: ejecución
FROM gcr.io/distroless/base-debian11


WORKDIR /

COPY --from=builder /app/app /app

EXPOSE 8084

CMD ["/app"]
