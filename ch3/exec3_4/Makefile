include ../../Makefile
EXEC = surface

build: main.go
	go build -o bin/$(EXEC) .

run: build
	./bin/$(EXEC) &

view: build
	open http://localhost:8000?width=640%26height=320%26minZColor=%230000FF%26maxZColor=%23FF0000
