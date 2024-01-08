build_amd64:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/main .

build_x86:
	mkdir -p bin/x86
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/x86/main .	