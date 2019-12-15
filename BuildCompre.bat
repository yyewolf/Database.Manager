set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
packr
go generate
go build -ldflags="-s -w" -ldflags -H=windowsgui
packr clean
D:\upx\upx.exe -9 -oDatabase.ManagerC.exe "%~dp0Database.Manager.exe"
del /f Database.Manager.exe
ren Database.ManagerC.exe Database.Manager.exe
PAUSE