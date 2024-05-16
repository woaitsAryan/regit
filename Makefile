.PHONY: setup

setup:
	@which go > /dev/null || (echo "needs go installed" && exit 1)
	go build -o regit
	chmod +x regit
	sudo mv regit /bin