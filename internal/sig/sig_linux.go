package sig

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignals(sigCh chan os.Signal){
	signal.Notify(sigCh,
		syscall.SIGABRT,
		syscall.SIGALRM,
		syscall.SIGFPE,
		syscall.SIGHUP,
		syscall.SIGILL,
		syscall.SIGINT,
		syscall.SIGIOT,
		//syscall.SIGKILL,
		syscall.SIGPIPE,
		syscall.SIGQUIT,
		syscall.SIGSEGV,
		syscall.SIGTERM,
		syscall.SIGTRAP)
}