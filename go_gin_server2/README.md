https://www.prakharsrivastav.com/posts/golang-small-docker-image/

https://www.youtube.com/watch?v=WPpw61vScIs


DOCKER

# START
    Build the docker image by running the below command from the project root.
    This will build and tag the docker image as basic:latest.

docker build -t basic:latest -f dockerfile-basic .

    Check the docker image size by running

docker images.

docker images
REPOSITORY   TAG               IMAGE ID       CREATED          SIZE
basic        latest            456f65ea2dfc   12 minutes ago   408MB
golang       1.17-alpine3.15   16b6bb19f174   3 days ago       315MB


# MULTISTAGE
docker build -t multistage:latest -f dockerfile-multistage .

REPOSITORY   TAG               IMAGE ID       CREATED          SIZE
multistage   latest            8e4262179789   2 minutes ago    14.9MB
basic        latest            456f65ea2dfc   59 minutes ago   408MB
golang       1.17-alpine3.15   16b6bb19f174   3 days ago       315MB
alpine       latest            c059bfaa849c   3 months ago     5.59MB


    We can further trim down the size of the docker image by providing additional build flags while building the native binary. This will strip down any debug information from the binary (this should be ok in some cases is not recommended for all scenarios) .



CURL

GET
curl http://localhost:8080/payments

curl http://localhost:8080/payments \
    --header "Content-Type: application/json" \
    --request "GET"

curl http://localhost:8080/payments/2


POST

curl http://localhost:8080/payments \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","invoice": "x467","currency": "EUR","amount": 49.99}'


Generate UUID
 go/src %  > for i in {1..3}; do UUID=`uuidgen`; echo $UUID; sleep 5;  done