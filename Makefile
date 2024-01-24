INSTALLPATH = ${HOME}/.local/bin

build:
	@sqlc generate -f sqlc/sqlc.yml
	@go build -o bin/clipstory cmd/clipstory/main.go
	@go build -o bin/clipstoryd cmd/clipstoryd/main.go

clean:
	@rm -rf bin/

install: build
	@cp -v bin/clipstory $(INSTALLPATH)
	@cp -v bin/clipstoryd $(INSTALLPATH)

remove:
	@rm -v $(INSTALLPATH)/clipstory
	@rm -v $(INSTALLPATH)/clipstoryd
