$env:GOOS='js'
$env:GOARCH='wasm'
go build -ldflags="-s -w" -v -o frontend/static/calculator.wasm ./wasm