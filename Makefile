# AUTHOR: Don Stringham
.DEFAULT_GOAL=clean

# VARIABLES
CC=go
BINARY_NAME=findlinks
SRC_DIR=cmd/non-concurrent
BLD_DIR=build

# DETERMINE OS
ifeq (${OS}, Windows_NT)
    BIN_POSTFIX=windows
else
    UNAME_S := $(shell uname -s)
    ifeq (${UNAME_S}, Linux)
        BIN_POSTFIX=linux
    endif
    ifeq (${UNAME_S}, Darwin)
        BIN_POSTFIX=darwin
    endif
endif

# TARGETS
build:
	-mkdir -p ${BLD_DIR}
	GOARCH=amd64 GOOS=darwin ${CC} build -o ${BLD_DIR}/${BINARY_NAME}-darwin ${SRC_DIR}/findlinks.go
	GOARCH=amd64 GOOS=linux ${CC} build -o ${BLD_DIR}/${BINARY_NAME}-linux ${SRC_DIR}/findlinks.go
	GOARCH=amd64 GOOS=windows ${CC} build -o ${BLD_DIR}/${BINARY_NAME}-windows ${SRC_DIR}/findlinks.go

run: build
	./${BLD_DIR}/${BINARY_NAME}-${BIN_POSTFIX} https://google.com

clean:
	go clean
	-rm -fr ${BLD_DIR}
	-rm -f `fzf -f .DS_Store`

test:
	${CC} test ${SRC_DIR}/...

.PHONY: build run test clean
