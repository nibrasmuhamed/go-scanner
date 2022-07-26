FROM golang:alpine3.15 as builder

LABEL maintainer="Go-scanner <https://github.com/nibrasmuhamed/go-scanner>"

WORKDIR /usr/src/go-scanner

COPY . .

RUN go get

RUN go build 

# USER go-scanner

ENTRYPOINT [ "./go-scanner" ]