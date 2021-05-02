.PHONY: build clean help

CMDS := $(foreach CMD,$(wildcard cmd/*),$(subst cmd/,,$(CMD)))

#? build: compile binaries in /cmd
build:
	@echo "Building commands..."
	@for CMD in $(CMDS); do \
  		echo "Building ./cmd/$$CMD"; \
  		go build -o bin/$$CMD -v ./cmd/$$CMD; \
  	done

#? clean: clean binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin/

#? help: display help
help: Makefile
	@printf "Available make targets:\n\n"
	@sed -n 's/^#?//p' $< | column -t -s ':' |  sed -e 's/^/ /'
