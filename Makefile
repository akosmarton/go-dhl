deps-clean:
	@rm -rf vendor/{src,pkg,bin}

deps-install: deps-clean
	@echo Installing dependencies...
	@go get -u -v github.com/constabulary/gb/...

deps-update:
	@echo Updating dependencies. This might take a minute ...
	@gb vendor update -all

deps-restore:
	@echo Restoring dependencies. This might take a minute ...
	gb vendor restore

test:
	@cd src && go test -v ./...

.PHONY: test
