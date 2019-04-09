package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/yuhu-tech/uche-data-sync/internal/collection/client"
	"github.com/yuhu-tech/uche-data-sync/internal/collection/config"
	"github.com/yuhu-tech/uche-data-sync/internal/common/log"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	conf, err := config.InitConfig()
	if err != nil {
		fmt.Printf("[Fatal] %v, Bye!", err)
		os.Exit(10)
	}

	log.Infof("Starting %s", os.Args[0])

	// quitChan is used to notify the goroutine to complete the current operation to exit.
	quitChan := make(chan bool)
	// doneChan is used to notify the main process that the goroutinue has completed the current operation.
	doneChan := make(chan bool)
	// fatalChan is used to nofity the main precess that the goroutinue encountered a fatal error and should exit right now.
	fatalChan := make(chan bool)
	// singalChan is used to notify the main process that the linux sent some signals what should be dealt with.
	signalChan := make(chan os.Signal, 1)

	go client.StartMQTTClient(conf, quitChan, doneChan, fatalChan)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case s := <-signalChan:
			quitChan <- true
			log.Infof("Captured %v. Prepare to exit the program gracefully", s)
			<-doneChan
			log.Info("Exiting...")
			os.Exit(0)
		case <-fatalChan:
			log.Error("The program crashed. eExiting...")
			os.Exit(0)
		}
	}
}
