FROM node:5

ADD ./web /var/www
WORKDIR /var/www

RUN npm install -g bower \
  && npm install -g grunt-cli \
  && rm -rf node_modules && rm -rf bower-components && rm -rf dist \
  && npm install \
  && bower install --allow-root=true --config.analytics=false --config.interactive=false \
  && grunt

FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ADD ./api /go/src/app
WORKDIR /go/src/app

RUN go get -d -v \
  && go get -u github.com/jteeuwen/go-bindata/... \
  && go-bindata -nocompress=true /var/www/web/dist/... \
  && go install -v 

RUN mkdir -p build/bitcannon \
  && cp config.json build/ \
  && cp config.json build/bitcannon/ \
  && go build -o build/bitcannon_bin

CMD ["build/bitcannon_bin"]