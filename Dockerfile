FROM golang:1.21.5 as builder
LABEL authors="Глеб"

WORKDIR /app
COPY go.mod .
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin .


FROM scratch as production

COPY --from=builder /app/bin /usr/bin/start

ENTRYPOINT ["usr/bin/start"]