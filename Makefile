.PHONY: all clean run win64

WIN_CC ?= x86_64-w64-mingw32-gcc

all:
	go generate ./...
	go build

win64:
	go generate ./...
	CGO_ENABLED=1 GOOS=windows CC=$(WIN_CC) go build -o blockgame.exe

run: all
	./blockgame

clean:
	rm -f *_gen.go blockgame blockgame.exe
