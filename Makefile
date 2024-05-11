.PHONY: build
.PHONY: setup

build:
	@which go > /dev/null || (echo "needs go installed" && exit 1)
	go build -o regit
	chmod +x regit