FROM golang:1.7

ADD . /go/src/app
WORKDIR /go/src/app

RUN ls -l /var \
  && go get -u github.com/jteeuwen/go-bindata/... \
  && go-bindata -nocompress=true web/...

RUN go get -d -v \
  && go install -v 

RUN mkdir -p build/bitcannon \
  && cp config.json build/ \
  && cp config.json build/bitcannon/ \
  && go build -o build/bitcannon_bin

CMD ["build/bitcannon_bin"]