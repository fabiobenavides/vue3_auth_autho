go build -o ../api_darwin_arm64
GOOS=darwin GOARCH=amd64 go build -o ../api_darwin_amd64
GOOS=linux GOARCH=amd64 go build -o ../api_linux_amd64
GOOS=windows GOARCH=amd64 go build -o ../api_windows_amd64

