FROM golang:alpine3.15

ENV GOPATH=/usr/go

LABEL maintainer="Go-scanner <https://github.com/nibrasmuhamed/go-scanner>"

WORKDIR /usr/src/go-scanner

COPY . .

RUN go get

RUN go install

ENTRYPOINT [ "/usr/go/bin/go-scanner" ]