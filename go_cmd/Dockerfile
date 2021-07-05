#https://docs.docker.com/language/golang/build-images/
# https://bitfieldconsulting.com/golang/docker-image

# tells Docker what base image to use for our application
# (i.e. our image will include all functionality from this base image and
# all subsequent coomands will build on top)
FROM golang:1.16-alpine
# Alpine is chosen for its small footprint compared to Ubuntu

# create a directory inside the image that we are building.
# this also instructs Docker to use this directory as the default destinatio for all subsequent commands. 
# This way we do not have to type out full file paths but can use relative paths based on this directory.
WORKDIR /app

# the very first thing you do once you’ve downloaded a project written in Go is to install
# the modules necessary to compile it.
# so need to get our go.mod and go.sum files copied into it, wuse the COPY command to do this.
# COPY command takes two parameters.
# - first parameter tells Docker what file you would like to copy into the image.
# - second parameter tells Docker where you want that file to be copied to.

# go.mod and go.sum are copied into our working directory /app which,
# owing to our use of WORKDIR, is the current directory (.) inside the image.
COPY go.mod .
#COPY go.sum .

#we have the module files inside the Docker image that were building
#so we can run the next command
RUN go mod download
# Go modules are now installed in /app andat this point,
# we have an image that is based on Go environment version 1.16
# and we have installed our dependencies.

# copy our source code into the image
# use a wildcard to copy all files with .go extension located in the current directory
# on the host (the directory where the Dockerfile is located)
# into the current directory inside the image.
COPY *.go .

# compile the application
RUN go build -o /go_cmd
# the result is a static application binary named go_cmd and
# located in the root of the filesystem of the image that we are building.
# the binary could have been put into any other place inside that image,
# the root directory has no special meaning in this regard,
# just convenient to use it to keep the file paths short for improved readability.

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
# EXPOSE 8080

# tells Docker what to execute when the image is used to start a container
CMD [ "/go_cmd run sleep 3" ]