https://docs.docker.com/language/golang/build-images/

https://github.com/olliefr/docker-gs-ping/blob/main/main.go

https://medium.com/codeshake/my-baby-steps-with-go-creating-and-dockerizing-a-rest-api-80522bc478cf

https://www.youtube.com/watch?v=WPpw61vScIs


## build
docker build -t mcf:latest -f dockerfile .

## run
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


## curl

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



while sleep 0.01; do curl http://localhost:8080/intros \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"prefix": "Hello World", "timestamp": "Jan 02, 2022 07:24:00 AM"}'; done

while sleep 0.1;
    do curl http://localhost:8080/payments \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","invoice": "x467","currency": "EUR","amount": 49.99}';
done


## hey

hey -n 8 -c 1  -m POST -D payment.json http://localhost:8080/payments