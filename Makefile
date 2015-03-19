build: build_web copy_web build_api
	@echo "BitCannon built to api/build"
deploy: build_web copy_web deploy_api package
	@echo "BitCannon releases zipped in the build folder."

build_web:
	@echo Building the web app...
	@cd web; \
	grunt
	@echo Finished building web.
build_api:
	@cd api; \
	make build
deploy_api:
	@cd api; \
	make deploy
copy_web:
	@echo Copying the web app to the api...
	@rm -rf api/web
	@cp -r web/dist api/web
	@touch api/web/.gitkeep
package:
	@cd api; \
	make package
