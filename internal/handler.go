package internal

import (
	"crypto/sha256"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"github.com/jmhobbs/parrotify-hd/pkg/parrot"
)

func decodeByExtension(body io.Reader, filename string) (image.Image, error) {
	split := strings.Split(strings.ToLower(filename), ".")
	extension := split[len(split)-1]
	switch extension {
	case "png":
		return png.Decode(body)
	case "jpeg":
		fallthrough
	case "jpg":
		return jpeg.Decode(body)
	case "gif":
		return gif.Decode(body)
	}
	return nil, fmt.Errorf("Unknown image filetype: %q", extension)
}

func decodeByContentType(body io.Reader, contentType string) (image.Image, error) {
	split := strings.Split(contentType, ";")
	switch split[0] {
	case "image/png":
		return png.Decode(body)
	case "image/jpeg":
		return jpeg.Decode(body)
	case "image/gif":
		return gif.Decode(body)
	}
	return nil, fmt.Errorf("Unknown image filetype: %q", split[0])
}

func MakeHandler(log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		overlaySrc := r.URL.Query().Get("src")

		if overlaySrc == "" {
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
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

		overlay, err := decodeByContentType(resp.Body, resp.Header.Get("content-type"))
		if err != nil {
			log.Error().Err(err).Str("src", overlaySrc).Msg("Unable to decode overlay image by content-type header")
			overlay, err = decodeByExtension(resp.Body, overlaySrc)
			if err != nil {
				log.Error().Err(err).Str("src", overlaySrc).Msg("Unable to decode overlay image by file name, giving up")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
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

		rotate, err := strconv.ParseFloat(r.URL.Query().Get("rotate"), 64)
		if err != nil {
			rotate = 0
		}

		gif, err := parrot.Overlay(overlay, int(scale), int(shiftX), int(shiftY), flip, rotate)
		if err != nil {
			log.Error().Err(err).Msg("Unable to compose gif")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		overlaySrcHash := sha256.New()
		overlaySrcHash.Write([]byte(overlaySrc))
		etag := fmt.Sprintf("%x_%d_%d_%d_%t", overlaySrcHash.Sum(nil), scale, shiftX, shiftY, flip)

		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Header().Set("ETag", fmt.Sprintf("%q", etag))
		w.Write(gif)
	}
}
