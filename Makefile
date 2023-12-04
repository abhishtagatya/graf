# Define variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

# Binary name
BINARY_NAME=graf

# Directories
OUT=./bin

all: clean build

build:
	$(GOBUILD) -o $(OUT)/$(BINARY_NAME) $(SRC)

clean:
	$(GOCLEAN)
	rm -f $(OUT)/$(BINARY_NAME)

example_dijkstra:
	$(GORUN) ./example/dijkstra.go

example_degree:
	$(GORUN) ./example/degree.go

example_export:
	$(GORUN) ./example/export.go

compute_container:
	$(GORUN) ./example/container.go "$(GRAPH)"

.PHONY: all build clean run example_dijkstra example_degree example_export compute_container
