build: build_web copy_web build_api
	echo "BitCannon built to api/build"
deploy: build_web copy_web deploy_api zip
	echo "BitCannon zipped to BitCannon.zip"

build_web:
	cd web; \
	grunt
build_api:
	cd api; \
	make build
deploy_api:
	cd api; \
	make deploy
copy_web:
	rm -rf api/web
	cp -r web/dist api/web
zip:
	cd api; \
	mv build bitcannon; \
	zip -r ../BitCannon.zip bitcannon; \
	mv bitcannon build
