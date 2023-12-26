FROM golang:1.21-alpine AS build

WORKDIR /usr/src/whmapper

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin -ldflags "-s" ./...

FROM scratch

COPY --from=build /usr/local/bin/whmserver /

CMD ["/whmserver"]