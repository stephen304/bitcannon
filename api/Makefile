build: clean
	go build -o build/bitcannon
deploy: clean
	gox -output="build/bitcannon_{{.OS}}_{{.Arch}}"
clean:
	rm -rf build
