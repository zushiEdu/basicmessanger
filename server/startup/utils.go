package startup

import (
	"main/databaseOperations"
	"os"
	"os/signal"
	"syscall"
)

func SetAwaitTermination() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		databaseOperations.CloseDB()
		os.Exit(1)
	}()
}
