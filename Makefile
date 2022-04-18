srcpath = ./src
files = $(srcpath)/main.go  $(srcpath)/command-parse.go $(srcpath)/commands.go $(srcpath)/event-handlers.go $(srcpath)/types-and-constants.go $(srcpath)/message-deleter.go $(srcpath)/format-logs.go


run:
	go run $(files)

build:
	go build -o ./builds/main.exe $(files)
