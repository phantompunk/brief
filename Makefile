.PHONY: all clean

binary   = brief
version  = 0.1.1
build	   = $(shell git rev-parse HEAD)
ldflags  = -ldflags "-X 'github.com/phantompunk/brief/command.version=$(version)'
ldflags += -X 'github.com/phantompunk/brief/command.build=$(build)'"

all:
	go build -o $(binary) $(ldflags)

clean:
	-rm $(binary)