package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	. "github.com/periaate/blume/hnet"
)

func main() {
	http.HandleFunc("POST /upload", func(w http.ResponseWriter, r *http.Request) {
		bar, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bar)
	})

	go func() { log.Fatalln(http.ListenAndServe("127.0.0.1:12719", nil)) }()

	response := URL("127.0.0.1:12719/upload").
		ToRequest(POST).
		FileAsBody("./main.go").
		WithHeaders(
			Content_Type.Tuple(Text_Plain),
		).
		Call().
		Assert(OK.Is).
		String(Println)

	if response.NetErr != nil {
		log.Fatalln(response.Error())
	}
	fmt.Println("Success!")
}
