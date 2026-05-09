FROM golang:1.24-alpine AS build

WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN go build -o /out/scrapem ./cmd/scrapem

FROM alpine:3.22

RUN apk add --no-cache ca-certificates
COPY --from=build /out/scrapem /usr/local/bin/scrapem
WORKDIR /vault

ENTRYPOINT ["/usr/local/bin/scrapem"]
