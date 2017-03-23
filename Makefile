.PHONY: default clean

CLI_NAME = thuglify
OS := $(shell uname)
VERSION ?= 1.0.0

# Path configuration #
CMD_DIR = cmd
BIN_DIR = bin

# target #

default: clean build_thugly

build_thugly: 
	@echo "Setup Thuglify"
ifeq ($(OS),Linux)
	mkdir -p build/linux
	@echo "Build Thuglify..."
	GOOS=linux  go build -ldflags "-X main.Version=$(VERSION)" -o build/linux/$(CLI_NAME) cmd/main.go
endif
ifeq ($(OS) ,Darwin)
	@echo "Build Thuglify..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o build/mac/$(CLI_NAME) cmd/main.go
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

install:
	echo "Install Thuglify, ${OS} version:= ${VERSION}"
ifeq ($(OS),Linux)
	mv build/linux/$(CLI_NAME) /usr/local/bin/$(CLI_NAME)
endif
ifeq ($(OS) ,Darwin)
	mv build/darwin/$(CLI_NAME) /usr/local/bin/$(CLI_NAME)
endif

clean:
	rm -rf build/*