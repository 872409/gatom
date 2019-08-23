package os

import (
	"os"
	"os/signal"
	"syscall"
)

func ListenExitSignal(exitHandle ...func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		// logger.Infoln(appName, "get a signal ", s.String())

		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			exitHandle[0]()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
