package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setUpDyncLogLevel(level string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Kill, os.Interrupt)
	select {
	case <-c:
		switch level {
		case "debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Hello")
	log.Debug().Msg("debug 1")

	go setUpDyncLogLevel("debug")
	time.Sleep(5 * time.Second)
	log.Debug().Msg("debug 2")

}
