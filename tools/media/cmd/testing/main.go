package main

import (
	"net/http"

	"github.com/periaate/blume/clog"
	"github.com/periaate/media"
)

func main() {
	http.HandleFunc("POST /resize", func(w http.ResponseWriter, r *http.Request) {
		img, err := media.ImageFromReader(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := media.MinImageSize(img, 768)
		media.Draw(img, res, true)
		media.FlushImage(w, res, 100)
		w.Header().Set("Content-Type", "image/jpeg")
	})

	clog.Info("listening on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
