GOCMD=go
GOBUILD=$(GOCMD) build -ldflags -s -v -a
BIN_BINARY_NAME=wxbot.exe

bot: export GOOS=windows
bot: export GOARCH=amd64
bot:
	@echo $(GOOS)
	@echo $(GOARCH)
	$(GOBUILD) -o $(BIN_BINARY_NAME) cmd/main.go
	mv $(BIN_BINARY_NAME) bin/

block: export GOOS=windows
block: export GOARCH=amd64
block: export BIN_BINARY_NAME=blockchain.exe
block:
	@echo $(GOOS)
	@echo $(GOARCH)
	$(GOBUILD) -o $(BIN_BINARY_NAME) xlsx/main.go
	mv $(BIN_BINARY_NAME) testdir/