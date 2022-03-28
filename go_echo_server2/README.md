https://docs.docker.com/language/golang/build-images/

https://github.com/olliefr/docker-gs-ping/blob/main/main.go

https://medium.com/codeshake/my-baby-steps-with-go-creating-and-dockerizing-a-rest-api-80522bc478cf

https://www.youtube.com/watch?v=WPpw61vScIs





Dockerfile      echo_server		scratch		    12.21 MB

Dockerfile_V1   echo_server		distroless  	27.14 MB

> docker run -p 80:8080 1695f4fbda5a
...
⇨ http server started on [::]:8080

localhost:8080
Safari can’t open the page "localhost:8080” because the server unexpectedly dropped the connection. This sometimes occurs when the server is busy. Wait for a few minutes, and then try again.

localhost:80/ping
{"Status":"OK"}


> docker run -p 100:10 1695f4fbda5a
...
⇨ http server started on [::]:8080

localhost:100/ping
Error, cannot open the page


> docker run -p 100:8080 1695f4fbda5a
...
⇨ http server started on [::]:8080

localhost:100/ping
{"Status":"OK"}


> docker run -p 100:80 1695f4fbda5a

localhost:100/ping
Error, cannot open the page

* docker build -t echo:v4 -f dockerfile3 .

## Build
# FROM golang:1.14-alpine AS build
FROM golang:1.17-alpine3.15 AS build

REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
echo         v4        908e9b3c3a73   13 hours ago   7.23MB
echo         v2        d78b3a6a1571   14 hours ago   12.2MB
echo         v3        e053efc4de9f   14 hours ago   12.2MB
echo         v1        1695f4fbda5a   4 days ago     12.2MB

