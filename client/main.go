package main

import (
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		keepAlive()
	}
}

func keepAlive() {
	defer log.Debug().Msg("connection closed")
	log.Debug().Msg("new connection")

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer log.Debug().Msg("stop interval")
	log.Debug().Msg("start interval")

	for range ticker.C {
		if _, err := conn.Write([]byte{1}); err != nil {
			return
		}
		log.Debug().Msg("keep alive")
	}
}
