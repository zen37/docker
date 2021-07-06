## Build
FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /go_echo_server

## Deploy
FROM scratch
COPY --from=build /go_echo_server /go_echo_server
ENTRYPOINT ["/go_echo_server"]