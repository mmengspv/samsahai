docker_registry ?= "quay.io"
docker_image_repo ?= "agoda-com/samsahai"
docker_image_tag ?= latest
#docker_user := ""
#docker_password := ""

package_name := github.com/agoda-com/samsahai
app_name := samsahai
output_path := ./out
go_ldflags ?= $(shell govvv -flags -pkg $(shell go list ./internal/samsahai))
golangci_lint_version := v1.15.0

.PHONY: init
init: tidy install-dep install

.PHONY: install-dep
install-dep:
	GO111MODULE=off go get github.com/ahmetb/govvv
	GO111MODULE=off go get golang.org/x/tools/cmd/goimports
	# install golangci-lint
	@curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | \
		sh -s -- -b $(shell go env GOPATH)/bin $(golangci_lint_version)

	@echo 'done!'

.PHONY: format
format:
	gofmt -w .
	goimports -w .

.PHONY: build
build: format
	GO111MODULE=on go build -ldflags="$(go_ldflags)" -o $(output_path)/$(app_name) cmd/main.go

.PHONY: install
install: build
	cp $(output_path)/$(app_name) $$GOPATH/bin/$(app_name)

.PHONY: print-flag
print-flag:
	@echo $(go_ldflags)

.PHONY: build-docker
build-docker:
	docker build -t $(docker_registry)/$(docker_image_repo):$(docker_image_tag) \
		--build-arg GO_LDFLAGS="$(go_ldflags)" \
		-f scripts/Dockerfile .

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: golangci-lint-check-version
golangci-lint-check-version:
	@if golangci-lint --version | grep "$(golangci_lint_version)" > /dev/null; then \
		echo; \
	else \
		echo "golangci-lint version mismatch"; \
		exit 1; \
	fi

.PHONY: coverage
coverage: format
	GO111MODULE=on golangci-lint run
	GO111MODULE=on go test -race -v `go list ./internal/... | grep -v cmd` -coverprofile=coverage.txt -covermode=atomic

.PHONY: cover-badge
cover-badge: coverage
	gopherbadger \
		-covercmd "go tool cover -func=coverage.txt"

.PHONY: coverage-html
coverage-html: coverage
	go tool cover -html=coverage.txt

.PHONY: docker-login-b64
docker-login-b64:
	#echo $(docker_password) | base64 --decode | docker login -u $(docker_user) $(docker_registry) --password-stdin

.PHONY: docker-login
docker-login:
	echo $(docker_password) | docker login -u $(docker_user) $(docker_registry) --password-stdin

.PHONY: docker-logout
docker-logout:
	docker logout $(docker_registry)

.PHONY: docker-push
docker-push:
	docker push $(docker_registry)/$(docker_image_repo):$(docker_image_tag)

.PHONY: docker-tag-n-push-latest
docker-tag-n-push-latest:
	@if [ "$(docker_image_tag)" != "latest" ]; then \
		docker tag $(docker_registry)/$(docker_image_repo):$(docker_image_tag) $(docker_registry)/$(docker_image_repo):latest; \
		docker push $(docker_registry)/$(docker_image_repo):latest; \
	fi
