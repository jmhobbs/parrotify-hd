package main

import (
	"image/png"
	"net/http"
	"os"
	"strconv"

	"github.com/jmhobbs/parrotify-hd/pkg/parrot"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main () {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	http.Handle("/", rootHandler())
	
	http.HandleFunc("/parrot.gif", func(w http.ResponseWriter, r *http.Request) {
		overlaySrc := r.URL.Query().Get("src")
		
		if overlaySrc == "" {
			w.Write(parrot.GifBytes)
			return
		}

		// todo: cap image size
		resp, err := http.Get(overlaySrc)
		if err != nil {
			log.Error().Err(err).Str("src", overlaySrc).Msg("Unable to get fetch overlay image")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// todo: handle other encodings
		overlay, err := png.Decode(resp.Body)
		if err != nil {
			log.Error().Err(err).Str("src", overlaySrc).Msg("Unable to decode overlay image")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		scale, err := strconv.ParseInt(r.URL.Query().Get("scale"), 10, 0)
		if err != nil {
			scale = 0
		}
		shiftX, err := strconv.ParseInt(r.URL.Query().Get("x"), 10, 0)
		if err != nil {
			shiftX = 0
		}
		shiftY, err := strconv.ParseInt(r.URL.Query().Get("y"), 10, 0)
		if err != nil {
			shiftY = 0
		}
		
		gif, err := parrot.Overlay(overlay, int(scale), int(shiftX), int(shiftY))
		if err != nil {
			log.Error().Err(err).Msg("Unable to compose gif")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(gif)
	})

	log.Info().Msg("Listening on 127.0.0.1:3333...")
	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if err != nil {
		log.Error().Err(err).Send()
	}

}