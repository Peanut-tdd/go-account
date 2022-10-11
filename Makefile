PROJECTNAME = go

## linux: 编译打包linux
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(RACE) -o main ./main.go

## win: 编译打包win
.PHONY: win
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(RACE) -o main-win.exe ./main.go

## mac: 编译打包mac
.PHONY: mac
mac:
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build $(RACE) -o main ./main.go
