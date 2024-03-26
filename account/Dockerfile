FROM golang:1.22.0-alpine3.19 as build

WORKDIR /go/src
COPY . .

RUN go mod tidy
RUN go build -o /bin/auth ./cmd/pulse/main.go

FROM alpine:3.19 as auth
COPY --from=build /bin/auth /bin/auth
CMD ["/bin/auth"]