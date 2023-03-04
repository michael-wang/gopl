build: main.go
	go build -o bin/server .

clean:
	rm -fr ./bin

run: build
	./bin/server

run3: build
	./bin/server --print

fetch := ../fetch/bin/fetch

$(fetch): ../fetch/main.go
	go build -o ../fetch/bin/fetch ../fetch

req: $(fetch)
	../fetch/bin/fetch http://localhost:8000
	../fetch/bin/fetch http://localhost:8000/help

req2: $(fetch)
	../fetch/bin/fetch http://localhost:8000
	../fetch/bin/fetch http://localhost:8000
	../fetch/bin/fetch http://localhost:8000/count
	../fetch/bin/fetch http://localhost:8000/count