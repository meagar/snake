.PHONY: html
html:
	GOOS=js GOARCH=wasm go build -ldflags "-w" -o docs/snake.wasm ./cmd/main.go
	cp -R assets docs/assets

