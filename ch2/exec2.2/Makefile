build: main.go
	go build -o bin/conv .

clean:
	rm -fr ./bin

run: build
	@echo "Using arguments:"
	@./bin/conv 100 40
	@echo "Using stdin:"
	@printf "100\n40\n" | ./bin/conv

