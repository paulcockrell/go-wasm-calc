compile:
	GOOS=js GOARCH=wasm go build -o ./webpack-starter/public/main.wasm main.go