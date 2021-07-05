##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /go_echo_server

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go_echo_server /go_echo_server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go_echo_server"]