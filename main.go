package main

import (
	"flag"
	"syscall"
	"github.com/tkanos/gonfig"
	"github.com/sevlyar/go-daemon"
	"os"
	"log"
	"github.com/sepuka/cryptoledger/structs"
	"time"
	"github.com/sepuka/cryptoledger/checker"
)

const (
	configPath = "./config.json"
	daemonName = "cryptoledger"
)

var (
	signal = flag.String("s", "", "send stop signal to the daemon")
	stop = make(chan struct{})

	config = structs.Configuration{}
	cntxt = &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{daemonName},
	}
)

func main() {
	readConfig()
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)

	if isDaemonFlagsPresent() {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %v", err)
		}
		daemon.SendCommands(d)

		return
	}

	child, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if child != nil {
		return
	} else {
		defer cntxt.Release()
	}

	log.Print("watcher daemon started")

	go mainLoop()

	err = daemon.ServeSignals()
	if err != nil {
		log.Println("Error:", err)
	}
	log.Println("daemon terminated.")
}

func isDaemonFlagsPresent() bool {
	return len(daemon.ActiveFlags()) > 0
}

func runWatchers() {
	checker.Ethereum(config.Ethereum, config)
}

func readConfig() {
	err := gonfig.GetConf(configPath, &config)
	if err != nil {
		log.Printf("Cannot read config: %v", err)
		os.Exit(1)
	}
}

func mainLoop() {
	for {
		select {
			case <-stop:
				break
			case <-time.After(time.Second*5):
				runWatchers()
		}
	}
}

func termHandler(sig os.Signal) error {
	log.Printf("Got '%v' signal, terminating...", sig)
	stop <- struct{}{}

	return daemon.ErrStop
}