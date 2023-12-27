FROM golang:1.21-alpine AS build

WORKDIR /usr/src/whmapper

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin -ldflags "-s" ./...

FROM alpine

RUN apk --no-cache add ca-certificates
COPY --from=build /usr/local/bin/whmserver /usr/local/bin/whmserver

CMD ["/usr/local/bin/whmserver"]