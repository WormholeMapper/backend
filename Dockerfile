FROM golang:1.21-alpine AS build

WORKDIR /usr/src/whmapper

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin ./...

FROM scratch

COPY --from=build /usr/local/bin/whmserver /

CMD ["/whmserver"]