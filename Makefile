CMD=go
SRC_PATH=src

all: clean build

build:
	${CMD} build -o exp ${SRC_PATH}/*

clean:
	rm -rf exp
