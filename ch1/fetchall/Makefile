build: main.go
	go build -o bin/fetchall .

clean:
	rm -fr ./bin ./out

run: build
	./bin/fetchall wikipedia.org

run-exec1.10: build
	mkdir -p ./out
	./bin/fetchall --save wikipedia.org

run-exec1.11: build
	cat exec1.11.input | xargs ./bin/fetchall