# BitCannon
A torrent site mirroring tool

## About
The goal of BitCannon is to provide the tools to easily aggregate the content of many torrent sites into an easily browse-able format.

BitCannon aims to be as user friendly as possible while still providing robustness and the features you would expect. We hope the average user will use BitCannon to keep personal bittorrent archives, but we strive to produce code that can stand up to running a public mirror as well.

## How to use: Simple Set-Up
Follow these steps to set up BitCannon for personal use:
* Download and install MongoDB for your platform
* Create the MongoDB data folder (`C:\data\db\` or similar)
* Download the latest release of BitCannon from the Releases tab
* Unzip the release somewhere that is convenient
* Run the mongod program to start the database (`C:\Program Files\MongoDB\bin\mongod.exe`)
* Run bitcannon for your platform from the release zip (`bitcannon.exe`)
* Open http://127.0.0.1:1337/ in your browser

You should see the BitCannon interface at this point

* Update/add torrents by downloading .gz torrent archives
* __Be sure to extract the txt file out of the gz file!__
* Drag the extracted text file onto the bitcannon.exe

## How to use: Building From Source

### MongoDB
* Install and run MongoDB from official packages

You must build the web first, as it gets embedded into the api binary.

### Web
* Install node (`sudo pacman -S nodejs`)
* Install bower and grunt with `sudo npm install -g grunt` and `sudo npm install -g grunt-cli`
* In `/web` type `npm install`, `bower install`, and `grunt`
* Check that the web built into the dist folder

### API
* Clone the repo
* Install go (`packer -S go-git`)
* Set $GOPATH (`export GOPATH=$HOME/.go`)
* Set $PATH (`export PATH="$PATH:$GOPATH/bin"`)
* Restart your terminal if you added these env vars to the startup script
* Copy `api/config.example.json` to `config.json`
* In the main folder, run `make build_api` to try to build
* If go complains about dependencies, get them with `go get <url>`
* Type `make build` in the main folder to compile the api into `build/`
* Optional: Cross compile for other platforms
  * Run `go get github.com/mitchellh/gox`
  * Build the toolchain with `gox -build-toolchain`
  * Compile with `make deploy` (Will make a zip containing all the binaries)
* Run `bitcannon` to run the api server
* Run `bitcannon <btArchive.txt>` to import torrents

## Progress
The early version of BitCannon aims to provide import functionality from bittorrent archives and a simple interface to browse and search your torrent database. Later versions may have more advanced features like auto updating torrent indexes and possibly more.

## License
This is MIT licensed, so do whatever you want with it. Just don't blame me for anything that happens.
