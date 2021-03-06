.PHONY: all clean build docker pki test-pki test-public test destroy-test-env

PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: clean build

clean:
	go clean -i ./...
	find . -name \*.out -type f -delete
	find . -name test-\*.log -type f -delete
	rm -f generate_pki generate_cert generate_cert.go

proto:
	protoc -I imap/ imap/node.proto --go_out=plugins=grpc:imap
	protoc -I comm/ comm/receiver.proto --go_out=plugins=grpc:comm

build:
	CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'

docker:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'
	docker build -t gopluto/pluto .

pki:
	go build crypto/generate_pki.go
	./generate_pki
	rm generate_pki

test-pki:
	go build crypto/generate_pki.go
	./generate_pki -pluto-config test-config.toml -rsa-bits 1024
	rm generate_pki

test-public:
	if [ ! -d "private" ]; then mkdir private; fi
	chmod 0700 private
	wget https://raw.githubusercontent.com/golang/go/master/src/crypto/tls/generate_cert.go
	go build generate_cert.go
	./generate_cert -ca -duration 2160h -host localhost,127.0.0.1,::1 -rsa-bits 1024
	mv cert.pem private/public-distributor-cert.pem && mv key.pem private/public-distributor-key.pem
	go clean
	rm -f generate_cert.go

test: destroy-test-env
	@echo "mode: atomic" > coverage.out;
	@echo ""
	@for PKG in $(PACKAGES); do \
		go test -v -race -coverprofile $${GOPATH}/src/$${PKG}/coverage-package.out -covermode=atomic $${PKG} || exit 1; \
		test ! -f $${GOPATH}/src/$${PKG}/coverage-package.out || (cat $${GOPATH}/src/$${PKG}/coverage-package.out | grep -v mode: | sort -r >> coverage.out); \
	done

destroy-test-env:
	if [ -d "private/Maildirs" ]; then rm -rf private/Maildirs; fi
	if [ -d "private/crdt-layers" ]; then rm -rf private/crdt-layers; fi
