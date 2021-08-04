package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felixge/httpsnoop"
	"github.com/p16n/pbdb/db"
	"github.com/spf13/viper"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Path[len("/"):]

	if k == "" {
		fmt.Fprint(w, "pbly is a link shortener")
		return
	}

	b, err := db.Get(k)
	if err != nil {
		log.Printf("Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(b) < 1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, string(b), http.StatusSeeOther)
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(mux, w, r)
		log.Printf(
			"%s %s (code=%d dt=%s written=%d)",
			r.Method,
			r.URL,
			m.Code,
			m.Duration,
			m.Written,
		)
	})

	port := fmt.Sprintf(":%s", viper.GetString("port"))

	log.Printf("Starting pbly server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, wrappedMux))
}
