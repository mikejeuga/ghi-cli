repo=$(shell basename "`pwd`")
gopher:
	@git init
	@go mod init github.com/mikejeuga/$(repo)
	@go mod tidy
	@mkdir cmd src models issues specifications testhelpers util
	@touch ./cmd/main.go
	@touch ./util/.gitkeep
	@touch .gitignore
	@touch Dockerfile
	@touch docker-compose.yml


t: test
test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v ./...
	@docker-compose down


ut: unit-test
unit-test:
	@go test -v -tags=unit ./...

at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -v -tags=acceptance ./...
	@docker-compose down

c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push
