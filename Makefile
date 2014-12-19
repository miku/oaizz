SHELL := /bin/bash
TARGETS = oaizz

# http://docs.travis-ci.com/user/languages/go/#Default-Test-Script
test:
	go get -d && go test -v

imports:
	goimports -w .

fmt:
	go fmt ./...

all: fmt test
	go build

install:
	go install

clean:
	go clean
	rm -f coverage.out
	rm -f $(TARGETS)
	rm -f oaizz-*.x86_64.rpm
	rm -f debian/oaizz*.deb
	rm -rf debian/oaizz/usr

cover:
	go get -d && go test -v	-coverprofile=coverage.out
	go tool cover -html=coverage.out

oaizz:
	go build -o oaizz cmd/oaizz/main.go

# ==== packaging

deb: $(TARGETS)
	mkdir -p debian/oaizz/usr/sbin
	cp $(TARGETS) debian/oaizz/usr/sbin
	cd debian && fakeroot dpkg-deb --build oaizz .

REPOPATH = /usr/share/nginx/html/repo/CentOS/6/x86_64

publish: rpm
	cp oaizz-*.rpm $(REPOPATH)
	createrepo $(REPOPATH)

rpm: $(TARGETS)
	mkdir -p $(HOME)/rpmbuild/{BUILD,SOURCES,SPECS,RPMS}
	cp ./packaging/oaizz.spec $(HOME)/rpmbuild/SPECS
	cp $(TARGETS) $(HOME)/rpmbuild/BUILD
	./packaging/buildrpm.sh oaizz
	cp $(HOME)/rpmbuild/RPMS/x86_64/oaizz*.rpm .
