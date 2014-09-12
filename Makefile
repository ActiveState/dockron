all:	fmt build
	@echo

build:
	docker build -t activestate/dockron .

run:
	docker run --rm -v /usr/local/bin/docker:/usr/bin/docker:ro -v /var/run/docker.sock:/var/run/docker.sock activestate/dockron "* * * * *" docker run ubuntu /bin/bash -c "echo Hello world"

test:
	docker run --rm --entrypoint=go activestate/dockron test -v ./... github.com/robfig/cron

fmt:
	gofmt -w .
