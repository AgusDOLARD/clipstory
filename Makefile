INSTALLPATH = ${HOME}/.local/bin

build:
	@go build -o bin/clipstory -ldflags "-s -w" cmd/clipstory/main.go

install: build
	@cp -v bin/clipstory $(INSTALLPATH)

remove:
	@rm -v $(INSTALLPATH)/clipstory
