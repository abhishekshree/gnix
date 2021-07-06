.PHONY: all

all: build comp-ls comp-cat comp-cp comp-dirname comp-echo comp-mkdir comp-pwd comp-sleep comp-touch comp-whoami comp-grep

.PHONY: clean

clean:
	rm -r build

build:
	mkdir -p build

comp-ls:
	go build -o build/ls ./ls/ls.go

comp-cat:
	go build -o build/cat ./cat/cat.go

comp-echo:
	go build -o build/echo ./echo/echo.go

comp-cp:
	go build -o build/cp ./cp/cp.go

comp-dirname:
	go build -o build/dirname ./dirname/dirname.go

comp-touch:
	go build -o build/touch ./touch/touch.go

comp-pwd:
	go build -o build/pwd ./pwd/pwd.go

comp-mkdir:
	go build -o build/mkdir ./mkdir/mkdir.go

comp-sleep:
	go build -o build/sleep ./sleep/sleep.go

comp-whoami:
	go build -o build/whoami ./whoami/whoami.go

comp-grep:
	go build -o build/grep ./grep/grep.go