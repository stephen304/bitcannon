build: clean prep config
	go build -o build/bitcannon_bin
deploy: clean prep config
	gox -output="build/bitcannon_{{.OS}}_{{.Arch}}"

clean:
	rm -rf build
prep:
	mkdir -p build/bitcannon
	go-bindata -nocompress=true web/...
config:
	cp config.json build/
	cp config.json build/bitcannon/

package:
	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_darwin_* build/bitcannon
	cd build; \
	zip -r BC_MacOSX_v$(VERSION).zip bitcannon

	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_freebsd_* build/bitcannon
	cd build; \
	zip -r BC_FreeBSD_v$(VERSION).zip bitcannon

	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_linux_* build/bitcannon
	cd build; \
	zip -r BC_Linux_v$(VERSION).zip bitcannon

	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_netbsd_* build/bitcannon
	cd build; \
	zip -r BC_NetBSD_v$(VERSION).zip bitcannon

	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_openbsd_* build/bitcannon
	cd build; \
	zip -r BC_OpenBSD_v$(VERSION).zip bitcannon

	rm -rf build/bitcannon/bitcannon*
	cp build/bitcannon_windows_* build/bitcannon
	cd build; \
	zip -r BC_Windows_v$(VERSION).zip bitcannon
