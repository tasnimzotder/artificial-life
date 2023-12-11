IMAGE_NAME = "tasnimzotder/artificial-life"

dkr-build:
	docker build -t $(IMAGE_NAME):latest .

dkr-run:
	docker run -it --rm -p 8080:8080 $(IMAGE_NAME):latest

wasm-build:
	env GOOS=js GOARCH=wasm go build -o artificial-life.wasm github.com/tasnimzotder/artificial-life

wasm-copy:
	cp $(go env GOROOT)/misc/wasm/wasm_exec.js .

mac-arm64-build:
	env GOOS=darwin GOARCH=arm64 go build -o bin/artificial-life-mac-arm64.bin

.PHONY: dkr-build dkr-run