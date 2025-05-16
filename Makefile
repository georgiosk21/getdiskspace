APP_NAME=getdiskspace
BUILD_DIR=build

all: linux darwin windows

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-linux ./main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-mac ./main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)-win.exe ./main.go

clean:
	rm -rf $(BUILD_DIR)

run:
	go run main.go usage --path /
