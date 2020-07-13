FROM golang:1.14

WORKDIR src/auth

RUN go get ./...
RUN go run .

COPY . .

CMD ["auth-app"]