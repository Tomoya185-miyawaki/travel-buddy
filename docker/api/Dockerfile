FROM golang:1.20.5-alpine

WORKDIR /var/www

RUN apk update && apk add git && go install github.com/cosmtrek/air@v1.40.4

CMD ["air", "-c", ".air.toml"]