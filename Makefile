.PHONY: addons

help:
	@echo "Management commands"
	@echo ""
	@echo "Usage:"
	@echo "  ## Root Commands"
	@echo "    make build                               Build project."
	@echo "    make test                                Run tests."
	@echo "    make clean                               Clean the directory tree of produced artifacts."
	@echo "    make lint                                Run static code analysis (lint)."
	@echo "    make format                              Run code formatter."
	@echo ""
	@echo "  ## Utility Commands"
	@echo "    make setup-git-hooks                     Command to setup git hooks."
	@echo ""

build:
	./make.sh "build"

test:
	go test -v ./... ./...

cover:
	go test -cover ./... 

clean:
	./make.sh "clean"

lint:
	./make.sh "lint"

format:
	./make.sh "format"

check-logs:
	./make.sh "checkLogs"

check-tests:
	./make.sh "checkTests"

setup-git-hooks:
	./make.sh "setupGitHooks"

goreleaser:
	goreleaser release --snapshot --skip-publish --rm-dist

tidy:
	go mod tidy

sec:
	gosec ./...

