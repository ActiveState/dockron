all:	fmt build
	@echo

build:
	docker build -t activestate/dockron .

run:
	docker run --rm activestate/dockron "* * * * *" /bin/bash -c "echo Hello world"

fmt:
	gofmt -w .
