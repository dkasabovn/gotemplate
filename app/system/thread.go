package system

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"template/app/system/log"
)

type globalSigHandler struct {
	pipe    chan os.Signal
	mutex   sync.Mutex
	canExit bool
}

var (
	globalSigHandlerOnce sync.Once
	globalSigHandlerInst *globalSigHandler
)

func GetSigHandler() *globalSigHandler {
	pipe := make(chan os.Signal, 1)
	signal.Notify(pipe, os.Interrupt)
	return &globalSigHandler{
		pipe:    pipe,
		mutex:   sync.Mutex{},
		canExit: false,
	}
}

func (s *globalSigHandler) SetExit(v bool) {
	s.mutex.Lock()
	s.canExit = v
	s.mutex.Unlock()
}

func (s *globalSigHandler) CanExit() (v bool) {
	s.mutex.Lock()
	v = s.canExit
	s.mutex.Unlock()
	return
}

func (s *globalSigHandler) Wait() {
	for sig := range s.pipe {
		log.Info("%T signal found", sig)
		for !s.CanExit() {
			time.Sleep(time.Second * 5)
		}
		return
	}
}
