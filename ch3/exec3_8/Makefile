include ../../Makefile
EXEC = mandelbrot

build: main.go
	go build -o bin/$(EXEC) .

run: build
	mkdir -p $(OUT)
	./bin/$(EXEC) -p c64 -z 64 -y -1 -o $(OUT)/c64_z4_y-1.png
	./bin/$(EXEC) -p c128 -z 64 -y -1 -o $(OUT)/c128_z4_y-1.png
	./bin/$(EXEC) -p bigfloat -z 64 -y -1 -o $(OUT)/bigfloat-z4_y-1.png
