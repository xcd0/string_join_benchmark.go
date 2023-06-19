version := $(shell go version)
uname   := $(shell uname -rpo)

all:
	@echo "\n## ${version} ${uname}\n"'```'  | tee -a readme.md
	@go test -bench . -benchmem     | tee -a readme.md
	@echo '```'                     | tee -a readme.md

