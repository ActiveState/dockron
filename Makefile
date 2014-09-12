all:	build
	@echo

build:
	docker build -t activestate/dockron .

run:
	docker run --rm activestate/dockron
