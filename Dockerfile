FROM golang:alpine

EXPOSE 6698/tcp

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

CMD ["ergonomadic", "run"]

RUN \
    apk add --update build-base git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/github.com/edmund-huber/ergonomadic
WORKDIR /go/src/github.com/edmund-huber/ergonomadic

COPY . /go/src/github.com/edmund-huber/ergonomadic

RUN go get -v -d
RUN go install -v
