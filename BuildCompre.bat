set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
packr
go generate
go build -ldflags="-s -w"
packr clean
del /f LoupGarouCompre.exe
D:\upx\upx.exe -9 -oLoupGarouCompre.exe "%~dp0Loup.Garou.exe"
PAUSE