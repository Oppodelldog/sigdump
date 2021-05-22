package sig

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterSignals(sigCh chan os.Signal) {
	signal.Notify(sigCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGILL,
		syscall.SIGTRAP,
		syscall.SIGABRT,
		syscall.SIGBUS,
		syscall.SIGFPE,
		//syscall.SIGKILL,
		syscall.SIGSEGV,
		syscall.SIGPIPE,
		syscall.SIGALRM,
		syscall.SIGTERM)
}
