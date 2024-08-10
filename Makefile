INSTALLPATH = ${HOME}/.local/bin

build:
	@go build -o bin/clipstory cmd/main.go

clean:
	@rm -rf bin/

install: build
	@cp -v bin/clipstory $(INSTALLPATH)

remove:
	@rm -v $(INSTALLPATH)/clipstory
