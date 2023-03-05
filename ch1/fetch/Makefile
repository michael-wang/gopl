build: main.go
	go build -o bin/fetch .

run: build
	./bin/fetch https://go.dev/

run-exec1.7: build
	./bin/fetch --copy https://go.dev/

run-exec1.8: build
	./bin/fetch --prefix go.dev/

run-exec1.9: build
	./bin/fetch --status https://go.dev/

clean:
	rm -fr ./bin
