FROM golang

WORKDIR . /go/src/auth

RUN go get ..
RUN go install github.com/AlekseiAnisimov/auth

ENTRYPOINT /go/bin/auth

EXPOSE 8080