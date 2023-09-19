include ../../Makefile
EXEC = surface

build: main.go
	go build -o bin/$(EXEC) .

run: build
	mkdir -p $(OUT)
	./bin/$(EXEC) > $(OUT)/out.svg
