SHELL := powershell.exe
.SHELLFLAGS := -Command


srcpath = ./src
files = $(srcpath)/main.go  $(srcpath)/command-parse.go $(srcpath)/commands.go $(srcpath)/event-handlers.go $(srcpath)/types-and-constants.go $(srcpath)/message-deleter.go $(srcpath)/format-logs.go


run:
	go run $(files)

build:
	go build -o ./builds/main.exe $(files)
	
buildwin:
	$$env:GOOS="windows"; $$env:GOARCH="amd64"; go build -o ./builds/win/main.exe $(files)

buildlinux:
	$$env:GOOS="linux"; $$env:GOARCH="amd64"; go build -o ./builds/linux/main.o $(files)


clear:
	rm -R ./builds
	rm ./logs.txt

	

