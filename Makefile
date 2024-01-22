.PHONY: os test build check-env-release release clean
.IGNORE: os test build clean

OS_NAME := $(shell uname -s | tr A-Z a-z)
os:
	echo $(OS_NAME)

NAME = palworld-save-backup
BASE_BUILD_DIR = build
BUILD_NAME = $(GOOS)-$(GOARCH)$(GOARM)
BUILD_DIR = $(BASE_BUILD_DIR)/$(BUILD_NAME)
VERSION? = dev

ifeq ($(GOOS),windows)
  ext=.exe
  archiveCmd=zip -9 -r $(NAME)-$(BUILD_NAME)-$(VERSION).zip $(BUILD_NAME)
else
  ext=
  archiveCmd=tar czpvf $(NAME)-$(BUILD_NAME)-$(VERSION).tar.gz $(BUILD_NAME)
endif

test:
	go test -race -v -bench=. ./...

build: clean test
	go build -mod=vendor

check-env-release:
	@ if [ "$(GOOS)" = "" ]; then \
		echo "Env variable GOOS not set"; \
		exit 1; \
   	fi
	@ if [ "$(GOARCH)" = "" ]; then \
		echo "Env variable GOARCH not set"; \
		exit 1; \
   	fi

release: check-env-release
	mkdir -p $(BUILD_DIR)
	cp LICENSE $(BUILD_DIR)/
	cp README.md $(BUILD_DIR)/
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod=vendor -ldflags "-s -w" -o $(BUILD_DIR)/$(NAME)$(ext) main.go
	cd $(BASE_BUILD_DIR) ; $(archiveCmd)

clean:
	go clean
	rm -vrf $(BASE_BUILD_DIR)
