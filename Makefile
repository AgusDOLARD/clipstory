INSTALLPATH = ${HOME}/.local/bin

build:
	@go build -o bin/clipstory cmd/clipstory/main.go
	@go build -o bin/clipstoryd cmd/clipstoryd/main.go

deamon: build
	@bin/clipstoryd

run: build
	@bin/clipstory

install: build
	@cp -v bin/clipstory $(INSTALLPATH)
	@cp -v bin/clipstoryd $(INSTALLPATH)

remove:
	@rm -v $(INSTALLPATH)/clipstory
	@rm -v $(INSTALLPATH)/clipstoryd
