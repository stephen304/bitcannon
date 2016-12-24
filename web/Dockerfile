FROM node:5

ADD . /var/www
WORKDIR /var/www

RUN npm install -g bower \
  && npm install -g grunt-cli \
  && rm -rf node_modules && rm -rf bower-components && rm -rf dist \
  && npm install \
  && bower install --allow-root=true --config.analytics=false --config.interactive=false

CMD ["grunt"]
