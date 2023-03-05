build: main.go
	go build -o bin/dup1 .

run: build
	cat test1.txt test2.txt test3.txt | ./bin/dup1

clean:
	rm -fr ./bin
