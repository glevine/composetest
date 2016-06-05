FROM golang

ADD . /go/src/github.com/glevine/composetest

RUN go install github.com/glevine/composetest

ENTRYPOINT /go/bin/composetest

EXPOSE 8080

