build:
	@go build -o bin/clipstory cmd/clipstory/main.go
	@go build -o bin/clipstoryd cmd/clipstoryd/main.go

deamon: build
	@./bin/clipstoryd

run: build
	@./bin/clipstory
