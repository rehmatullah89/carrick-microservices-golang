package logger

import (
	"log"
	"sync"
)

type loggerSingleton struct {
}

var (
	loggerOnce     sync.Once
	loggerInstance *loggerSingleton
)

func GetLoggerInstance() *loggerSingleton {
	if loggerInstance == nil {
		loggerOnce.Do(
			func() {
				loggerInstance = &loggerSingleton{}
			})
	}

	return loggerInstance
}

func (s *loggerSingleton) Info(message interface{}) {
	log.Println("INFO:", message)
}

func (s *loggerSingleton) Warning(message interface{}) {
	log.Println("WARN:", message)
}

func (s *loggerSingleton) Error(message interface{}) {
	log.Println("ERROR:", message)
}

func (s *loggerSingleton) Fatal(message interface{}) {
	log.Fatalln("FATAL:", message)
}

func (s *loggerSingleton) Panic(message interface{}) {
	log.Panicln("PANIC:", message)
}

func (s *loggerSingleton) Debug(message interface{}) {
	log.Println("DEBUG:", message)
}
