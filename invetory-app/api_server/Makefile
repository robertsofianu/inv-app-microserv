VERSION_FILE := version.txt
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S)

# Dynamically read the current version from version.txt
VERSION := $(shell cat $(VERSION_FILE))

build:
	@echo "Current version: $(VERSION)"
	@read -p "Enter update type (MJ for major, default is minor): " UPDATE; \
	if [ "$$UPDATE" = "MJ" ]; then \
		echo $(VERSION) | awk -F. '{print $$1+1".0.0"}' > $(VERSION_FILE); \
	else \
		echo $(VERSION) | awk -F. '{print $$1"."$$2+1".0"}' > $(VERSION_FILE); \
	fi
	@echo "Version updated to $$(cat $(VERSION_FILE))"
	@echo "Building application with version $(VERSION)..."
	go build -ldflags "-X main.AppVersion=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o inventory-app ./CRUD/login_server.go
	@echo "Build completed."

run:
	@echo "Running application..."
	./inventory-app