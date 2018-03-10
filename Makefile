GOCMD=go
GOGET=$(GOCMD) get

deps:
	$(GOGET) github.com/nlopes/slack
	$(GOGET) github.com/tkanos/gonfig
	$(GOGET) github.com/sevlyar/go-daemon
