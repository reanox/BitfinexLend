FROM golang:1.12-alpine

RUN apk add make git

ADD . /bitfinex-lend-server
WORKDIR /bitfinex-lend-server

RUN git config --global credential.helper "store --file `pwd`/.git-credentials"
CMD ["go", "run", "cmd/bitfinex-lend-server/main.go"]