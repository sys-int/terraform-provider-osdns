TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=sys-int
NAMESPACE=test
NAME=opnsense
BINARY=terraform-provider-${NAME}
VERSION=0.3.4
OS_ARCH=linux_amd64
#OS_ARCH=windows_amd64

default: install

build:
	go build -o ${BINARY}

build-debug:
	go build -gcflags="all=-N -l" -o ${BINARY}

debug: build-debug
	dlv exec --accept-multiclient --continue --headless ./${BINARY} -- -debug

release:
	goreleaser release --rm-dist --snapshot --skip-publish  --skip-sign

install: build
	rm -rf ~/.terraform.d/plugins/${HOSTNAME}/${NAME}
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/
#	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
#	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m