SOURCEDIR=.

GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOBUILDVAR=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BINARY?=gogl
BINARY_PATH=$(SOURCEDIR)/main.go

.DEFAULT_GOAL: $(BINARY)

all: clean build

run: build
	./$(BINARY);
	@make clean

build: 
	$(GOBUILD) -o ${BINARY} $(BINARY_PATH)

download:
	$(GOMOD) download

clean: 
	rm -f $(BINARY)
