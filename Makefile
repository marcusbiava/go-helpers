name := go-helpers
module := github.com/marcusbiava/$(name)

hello:
	echo "Hello from $(module)"
	
go-mod-init:
	go mod init $(module)

init: hello go-mod-init

release:
	$(eval v := $(shell git describe --tags --abbrev=0 | sed -Ee 's/^v|-.*//'))
ifeq ($(bump), major)
	$(eval f := 1)
else ifeq ($(bump), minor)
	$(eval f := 2)
else
	$(eval f := 3)
endif
	$(eval newTag := $(shell echo $(v) | awk -F. -v OFS=. -v f=$(f) '{ $$f++ } 1'))
	@echo "v$(newTag)"
	@git tag "v$(newTag)"
	@git push --tags