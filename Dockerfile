FROM golang:1.22.1 AS build

WORKDIR /app

COPY ./ ./

RUN go build -o /byteApi

EXPOSE 8080

CMD [ "/byteApi" ]