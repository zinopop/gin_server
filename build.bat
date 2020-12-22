SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o bin/dingTalk_linux main.go

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o bin/dingTalk_windows.exe main.go