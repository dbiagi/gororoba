package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type HttpResponse struct {
	StatusCode  int
	Body        any
	ContentType string
	Headers     map[string]string
}

type HttpHandlerFunc func(w http.ResponseWriter, r *http.Request) HttpResponse

func HandleRequest(fn HttpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fn(w, r)

		if response.StatusCode != 0 {
			w.WriteHeader(response.StatusCode)
		}

		if response.Headers != nil {
			for key, value := range response.Headers {
				w.Header().Set(key, value)
			}
		}

		if response.Headers["Content-Type"] == "" {
			w.Header().Set("Content-Type", "application/json")
		}

		var bytes []byte
		var marshalingError error
		if response.Body != nil {
			bytes, marshalingError = json.Marshal(response.Body)

			if marshalingError != nil {
				slog.Error(
					fmt.Sprintf("Error marshalling response: %v", marshalingError.Error()),
					slog.String("error", marshalingError.Error()),
				)
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Write(bytes)
		}
	}
}
