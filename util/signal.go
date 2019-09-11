package util

import (
	"os"
	"os/signal"
	"syscall"
)

var handled = false

func HandleSignal(handler func(sig os.Signal, exist bool)) {
	if handled {
		return
	}
	handled = true

	chSignal := make(chan os.Signal, 1)

	// ctrl + c 	SIGINT 		强制进程结束
	// ctrl + z 	SIGTSTP 	任务中断，进程挂起
	// ctrl + \	 	SIGQUIT 	进程结束 和 dump core
	// ctrl + d	  		 		EOF
	//  			SIGHUP  	终止收到该信号的进程。若程序中没有捕捉该信号，当收到该信号时，进程就会退出（常用于 重启、重新加载进程）
	signal.Notify(chSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for {
		if sig, ok := <-chSignal; ok {
			var isExitSignal = false

			switch sig {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				isExitSignal = true
			default:
				isExitSignal = false
			}

			if handler != nil {
				handler(sig, isExitSignal)
			}

			if isExitSignal {
				return
			}

		}

	}
}
