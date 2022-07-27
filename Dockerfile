FROM golang:alpine3.15

# RUN addgroup -g 999 appuser && \
#     adduser -r -u 999 -g appuser appuser
# USER appuser
ENV GOPATH=/usr/go

LABEL maintainer="Go-scanner <https://github.com/nibrasmuhamed/go-scanner>"

WORKDIR /usr/src/go-scanner

COPY . .

RUN go get

RUN go install

# USER go-scanner

ENTRYPOINT [ "/usr/go/bin/go-scanner" ]