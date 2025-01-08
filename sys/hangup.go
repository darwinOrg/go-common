package dgsys

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func HangupApplication() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Kill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	sin := <-s
	log.Printf("application stoped，signal: %s", sin.String())
}
