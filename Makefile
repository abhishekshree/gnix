GOC=go build 

.PHONY: all

all: build ls cat cp dirname echo mkdir pwd sleep touch whoami

.PHONY: clean

clean:
	rm -r build

build:
	mkdir -p build

ls:
	${GOC}  -o build/ls ls.go

cat:
	${GOC} -o build/cat cat.go

echo:
	${GOC} -o build/echo echo.go

cp:
	${GOC} -o build/cp cp.go

dirname:
	${GOC} -o build/dirname dirname.go

touch:
	${GOC} -o build/touch touch.go

pwd:
	${GOC} -o build/pwd pwd.go

mkdir:
	${GOC} -o build/mkdir mkdir.go

sleep:
	${GOC} -o build/sleep sleep.go

whoami:
	${GOC} -o build/whoami whoami.go