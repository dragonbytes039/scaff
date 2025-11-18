APP_NAME := scaff
OUTPUT_DIR := ./bin
DEV_DIR := ./dev



$(OUTPUT_DIR):
	@mkdir -p $(OUTPUT_DIR)

##build: build the aplication 
.PHONY: build
build: 
	go mod tidy
	GOARCH=amd64 GOOS=darwin go build -o $(OUTPUT_DIR)/$(APP_NAME)-darwin ./
	GOARCH=amd64 GOOS=linux go build -o $(OUTPUT_DIR)/$(APP_NAME)-linux ./
	GOARCH=amd64 GOOS=windows go build -o $(OUTPUT_DIR)/$(APP_NAME)-windows.exe ./

.PHONY: build-cmd
build-cmd: 
	go mod tidy
	set GOARCH=amd64&&set GOOS=darwin&&go build -o $(OUTPUT_DIR)/$(APP_NAME)-darwin ./
	set GOARCH=amd64&&set GOOS=linux&&go build -o $(OUTPUT_DIR)/$(APP_NAME)-linux ./
	set GOARCH=amd64&&set GOOS=windows&&go build -o $(OUTPUT_DIR)/$(APP_NAME)-windows.exe ./


## windows dev 
.PHONY: build-w
build-w:
	go build -o $(DEV_DIR)/$(APP_NAME).exe

.PHONY: clean
clean: 
	rmdir /S /Q dev
	rmdir /S /Q project

.PHONY: run-windows
run-windows:
	./dev/scaff.exe create project

.PHONY: dev-w
dev-w: build-w run-windows clean

## windows dev