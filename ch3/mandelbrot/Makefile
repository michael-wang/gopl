include ../../Makefile
EXEC = mandelbrot

build: main.go
	go build -o bin/$(EXEC) .

run: build
	mkdir -p $(OUT)
	./bin/$(EXEC) > $(OUT)/out.png
