PKG=github.com/link-u/the-simplest-fastcgi-server-in-this-world

all: .bin/srv .bin/srv-linux

.bin:
	mkdir .bin

.bin/srv: main.go
	go build -o $@ $(PKG)

.bin/srv-linux: main.go
	GOOS=linux GOARCH=amd64 go build -o $@ $(PKG)

clean:
	rm -Rfv .bin

.PHONY: all clean
