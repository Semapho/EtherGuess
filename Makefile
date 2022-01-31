# Binary name
BINARY=ether-guess
VERSION=0.1
LDFLAGS='-w -s'

build:
	# Build
	rm -f ./${BINARY}
	go build -race -o ${BINARY} cmd/*.go

mac:
	# Build for mac
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	upx ./${BINARY}

mac-arm:
	# Build for mac-arm
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	upx ./${BINARY}

linux:
	# Build for linux
	rm -f ./${BINARY}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} -ldflags ${LDFLAGS} cmd/*.go
	upx ./${BINARY}

windows:
	# Build for windows
	rm -f ./${BINARY}.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe -ldflags ${LDFLAGS} cmd/*.go
	upx ./${BINARY}.exe

clean:
	# Clean projects
	rm -f ${BINARY}
	rm -f ${BINARY}.exe
	rm -f release.zip
	rm -rf release/

.PHONY: build mac mac-arm linux windows clean
