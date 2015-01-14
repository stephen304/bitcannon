build: clean prep config
	go build -o build/bitcannon
deploy: clean prep config
	gox -output="build/bitcannon_{{.OS}}_{{.Arch}}"

clean:
	rm -rf build
prep:
	mkdir build
	go-bindata -nocompress=true web/...
config:
	cp config.json build/
