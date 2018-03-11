GOCMD=go
GOGET=$(GOCMD) get

deps:
	$(GOGET) github.com/nlopes/slack
	$(GOGET) github.com/tkanos/gonfig
	$(GOGET) github.com/sevlyar/go-daemon

test:
	$(GOCMD) test -v ./...

daemon_start:
	$(GOCMD) run main.go

daemon_start_log: daemon_start
	tail -f log

daemon_stop:
	$(GOCMD) run main.go -s stop