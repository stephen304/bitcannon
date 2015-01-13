build: clean prep configure
	go build -o build/bitcannon
deploy: clean prep configure
	gox -output="build/bitcannon_{{.OS}}_{{.Arch}}"

clean:
	rm -rf build
prep:
	mkdir build
config:
	cp config.json build/
