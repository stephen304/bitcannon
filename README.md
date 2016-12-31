# BitCannon
A torrent site mirroring tool

## About
The goal of BitCannon is to provide the tools to easily aggregate the content of many torrent sites and DHT network into an easily browse-able format.

BitCannon aims to be as user friendly as possible while still providing robustness and the features you would expect. We hope the average user will use BitCannon to keep personal bittorrent archives, but we strive to produce code that can stand up to running a public mirror as well.

## Project Mirrors
This project is based from :
* [GitHub](https://github.com/Stephen304/bitcannon)
* [BitBucket](https://bitbucket.org/Stephen304/bitcannon)
* [Google Code](https://code.google.com/p/bitcannon/)

## How to use: Using Docker

Launch the docker-compose on root directory :

```
  docker-compose up -d
```

Then, you'll have the front available on localhost :

```
  firefox http://localhost
```

The API will be available at http://localhost:8000 so set it in settings.

You're done.

## How to use: Building From Source

> If you are not a programmer or do not wish to install this long list of things, use the instructions on the wiki instead!
* NodeJS
* Grunt
* Bower
* Golang
* Golang Dependencies

__(Note: These building instructions may get out of date from time to time due to code changes. If you just want to use BitCannon, you should use the Wiki instructions instead.)__

### MongoDB
* Install and run MongoDB from official packages

You must build the web first, as it gets embedded into the api binary.

### Web
* Install node (`sudo pacman -S nodejs`)
* Install bower and grunt with `sudo npm install -g grunt` and `sudo npm install -g grunt-cli`
* In `/web` type `npm install`, `bower install`, and `grunt`
* Check that the web built into the dist folder

> If grunt fails with errors, you may have not installed it properly. The NodeJS and Grunt guys probably know more about it than I do

### API
* Clone the repo
* Install go (`packer -S go-git`)
* Set $GOPATH (`export GOPATH=$HOME/.go`)
* Set $PATH (`export PATH="$PATH:$GOPATH/bin"`)
* Restart your terminal if you added these env vars to the startup script

> Go can be hard to install without nice official packages. If go spits errors, try googling them a bit before opening an issue. It may not be specific to this project.

* Copy `api/config.example.json` to `config.json`
* In the main folder, run `make build_api` to try to build
* If go complains about dependencies, get them with `go get <url>`

Once you have all of the dependencies, it will build into the api/build folder.

* Run `bitcannon` to run the server
* Run `bitcannon <btArchive.txt>` to import torrents

#### Extra things and tips
* If you edit the web app, typing `make build` in the main folder will recompile both the web and api into `api/build`
* If you only edited the api folder, use `make build_api` to avoid recompiling the web
* Optional: Cross compile for other platforms (Your go installation must be from the source or it will spit errors)
  * Run `go get github.com/mitchellh/gox`
  * Build the toolchain with `gox -build-toolchain`
  * Compile with `make deploy` (Will make a zip containing all the binaries)

## Progress
The early version of BitCannon aims to provide import functionality from bittorrent archives and a simple interface to browse and search your torrent database. Later versions may have more advanced features like auto updating torrent indexes and possibly more.

## Works great with DHTBay

Actually I've made this fork to work properly with the DHTBay project I've made, it is totally compatible, and I'll soon share with you my docker-compose file to make them work together.

Meanwhile you can access the project at https://github.com/FlyersWeb/dhtbay.

## License
This is MIT licensed, so do whatever you want with it. Just don't blame me for anything that happens.
