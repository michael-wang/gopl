build: main.go
	go build -o bin/lissajour .

run: build
	mkdir -p out
	bin/lissajour > out/lissajour.gif
	open -a Google\ Chrome out/lissajour.gif

run-exec1.5: build
	mkdir -p out
	bin/lissajour -green > out/lissajour.gif
	open -a Google\ Chrome out/lissajour.gif

run-exec1.6: build
	mkdir -p out
	bin/lissajour -rand > out/lissajour.gif
	open -a Google\ Chrome out/lissajour.gif

clean:
	rm -fr ./bin ./out
