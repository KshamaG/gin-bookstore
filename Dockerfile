FROM golang:1.15.11-alpine3.13
RUN apk add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8085
CMD ["/app/main"]