ARG GO_VERSION=1.23

FROM  --platform=$BUILDPLATFORM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0

WORKDIR /usr/src/app

COPY go.mod ./ 
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

FROM scratch

USER 65534
WORKDIR /usr/local/bin/

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/app /app

ENTRYPOINT ["/app"]
