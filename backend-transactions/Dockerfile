FROM golang:alpine AS main

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.* .

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build -o /go/bin/main cmd/server/main.go

FROM scratch

WORKDIR /app

COPY --from=main /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=main /go/bin/main /app

ENTRYPOINT ["/app/main"]