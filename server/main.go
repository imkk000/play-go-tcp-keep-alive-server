package main

import (
	"context"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l, err := net.Listen("tcp", "127.0.1:9000")
	if err != nil {
		log.Fatal().Err(err).Msg("listener")
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(ctx, conn)
	}
}

func handler(_ context.Context, conn net.Conn) {
	defer conn.Close()
	defer log.Info().Msg("connection closed")
	log.Info().
		Str("addr", conn.RemoteAddr().String()).
		Msg("new connection")

	bs := make([]byte, 1)
	for {
		if _, err := conn.Read(bs); err != nil {
			return
		}
		if _, err := conn.Write([]byte{2}); err != nil {
			return
		}
	}
}
