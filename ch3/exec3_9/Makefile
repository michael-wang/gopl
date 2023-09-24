include ../../Makefile
EXEC = mandelbrot

build: main.go
	go build -o bin/$(EXEC) .

run: build
	mkdir -p $(OUT)
	./bin/$(EXEC)

view: build
	open http://localhost:8000/?x=-0.1016&y=-0.9513&z=4096