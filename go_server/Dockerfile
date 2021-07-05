##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /go_server

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go_server /go_server

EXPOSE 8888

USER nonroot:nonroot

ENTRYPOINT ["/go_server"]