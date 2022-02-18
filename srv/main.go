package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jmhobbs/parrotify-hd/internal"
)

func main () {
	var port *int = flag.Int("PORT", 3333, "Port to listen on.")
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(logSink())

	http.Handle("/", rootHandler())
	http.HandleFunc("/parrot.gif", internal.MakeHandler(log.Logger))

	log.Info().Msgf("Listening on http://127.0.0.1:%d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Error().Err(err).Send()
	}

}