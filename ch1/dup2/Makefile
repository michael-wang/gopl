build: main.go
	go build -o bin/dup2 .

run-stdin: build
	cat test1.txt test2.txt test3.txt | bin/dup2

run-files: build
	bin/dup2 test1.txt test2.txt test3.txt

clean:
	rm -fr ./bin
