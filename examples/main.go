package main

import (
	"github.com/1gorAn/logger/logger"
)

func main() {
	log := logger.NewLogger(&logger.Config{
		Level:    "debug",   // доступные уровни: debug, info, warn, error, dpanic, panic, fatal
		Encoding: "console", // или "json"
	})

	log.InitLogger()

	log.Info("Logger инициализирован")
	log.Debugf("Debug с параметром: %s", "example")
}
