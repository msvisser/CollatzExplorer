BINARY := collatz
STATIC_FOLDER := ./static
STATIC_FILES := $(shell find ${STATIC_FOLDER} -type f -name '*.css' -or -name '*.js')
TEMPLATES_FOLDER := ./templates
TEMPLATES_FILES := $(shell find ${TEMPLATES_FOLDER} -type f)

.PHONY: all
all: ${BINARY}

.PHONY: clean
clean:
	rm -f ${BINARY} bindata.go

bindata.go: ${STATIC_FILES} ${TEMPLATES_FILES}
	go-bindata -nomemcopy -fs ${STATIC_FILES} ${TEMPLATES_FILES}

${BINARY}: collatz.go bindata.go
	GOARM=6 GOARCH=arm CGO_ENABLED=0 GOOS=linux \
	go build -a -tags netgo -ldflags '-w' -o ${BINARY} $^
