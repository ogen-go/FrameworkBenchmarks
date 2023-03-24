package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/go-faster/errors"

	"ogen/app/internal/api"
	"ogen/app/internal/oas"
)

func run(ctx context.Context) error {
	bindHost := flag.String("bind", ":8080", "set bind host")
	flag.Parse()

	server, err := oas.NewServer(api.Server{})
	if err != nil {
		return errors.Wrap(err, "failed to create server")
	}
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", "ogen")
			server.ServeHTTP(w, r)
		}),
		Addr: *bindHost,
	}

	return s.ListenAndServe()
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		panic(err)
	}
}
