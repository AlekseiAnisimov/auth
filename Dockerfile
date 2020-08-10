FROM golang

WORKDIR /app

COPY ./ /app

RUN go get ./...
RUN go install

ENTRYPOINT /go/bin/auth

EXPOSE 8080