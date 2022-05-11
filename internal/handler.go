package internal

import (
	"image/png"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"

	"github.com/jmhobbs/parrotify-hd/pkg/parrot"
)

func MakeHandler(log zerolog.Logger) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
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

		flip := r.URL.Query().Get("flip") == "true"
		
		gif, err := parrot.Overlay(overlay, int(scale), int(shiftX), int(shiftY), flip)
		if err != nil {
			log.Error().Err(err).Msg("Unable to compose gif")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(gif)
	}
}