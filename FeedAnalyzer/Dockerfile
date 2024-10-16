ARG GO_VERSION=1.23

FROM  --platform=$BUILDPLATFORM golang:${GO_VERSION} AS builder

WORKDIR /usr/src/app

COPY go.mod ./ 
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

FROM scratch

WORKDIR /usr/src/app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

USER nobody

ENTRYPOINT [ "app" ]
