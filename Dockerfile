FROM golang:alpine

EXPOSE 6697/tcp

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

CMD ["ircd", "run"]

RUN \
    apk add --update build-base git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/github.com/prologic/ircd
WORKDIR /go/src/github.com/prologic/ircd

COPY . /go/src/github.com/prologic/ircd

RUN go get -v -d
RUN go install -v
