.PHONY: clean

# Path configuration #
CMD_DIR = cmd
BIN_DIR = bin

# target #

default: clean build_thugly

build_thugly: 
	cd $(CMD_DIR); \
	go build -o thuglify; \
	cd ..; \
	mv $(CMD_DIR)/thuglify $(BIN_DIR)/ 

clean:
	rm -rf bin/
	mkdir bin