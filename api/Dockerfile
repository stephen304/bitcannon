FROM golang:alpine

COPY . /go/src/app

WORKDIR /go/src/app

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh \
  && go get -d -v \
  && go get -u github.com/jteeuwen/go-bindata/... \
  && go-bindata -nocompress=true web/... \
  && go install -v 

RUN mkdir -p build/bitcannon \
  && cp config.json build/ \
  && cp config.json build/bitcannon/ \
  && go build -o build/bitcannon_bin

CMD ["build/bitcannon_bin"]