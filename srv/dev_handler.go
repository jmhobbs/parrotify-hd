// +build !production

package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

func rootHandler() http.Handler {
	log.Info().Msg("[DEV] Proxing / to localhost:3000")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyURL, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		proxyURL.Scheme = "http"
		proxyURL.Host = "localhost:3000"

		proxyRequest, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		for header, values := range r.Header {
			for _, value := range values {
				proxyRequest.Header.Add(header, value)
			}
		}
		w.Header().Set("Host", "localhost:3000")

		response, err := http.DefaultClient.Do(proxyRequest)
		if err != nil {
			log.Error().Err(err).Msg("Error on request to proxy")
			if strings.Contains(err.Error(), "connection refused") {
				log.Error().Msg("Are you sure the frontend is running?")
			}
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for header, values := range response.Header {
			for _, value := range values {
				w.Header().Add(header, value)
			}
		}
		w.WriteHeader(response.StatusCode)

		_, err = io.Copy(w, response.Body)
		defer response.Body.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error copying from proxy request.")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})
}