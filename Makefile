.PHONY: all clean run

all:
	go generate ./...
	go build

run: all
	./blockgame

clean:
	rm -f *_gen.go blockgame
