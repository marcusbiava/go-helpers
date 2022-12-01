name := go-helpers
module := github.com/marcusbiava/$(name)

hello:
	echo "Hello from $(module)"
	
go-mod-init:
	go mod init $(module)

init: hello go-mod-init	