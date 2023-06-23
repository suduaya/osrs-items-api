package main

import (
	"fmt"
	"net/http"
	"osrs-items-api/internal/items"
	"osrs-items-api/pkg/oldschoolrs"
	"time"

	"github.com/gorilla/mux"
	httpswagger "github.com/swaggo/http-swagger"

	// This import is necessary for swagger documentation.
	_ "osrs-items-api/api"
)

func main() {
	// Avoid "404 page not found".
	router := mux.NewRouter()

	c := make(chan struct{}, 100) // max requests example

	router.Use(func(next http.Handler) http.Handler {
		// Limiting the degree of concurrency.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Counting semaphore using a buffered channel.
			select {
			case c <- struct{}{}:
				defer func() { <-c }()

				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			default:
				w.WriteHeader(http.StatusTooManyRequests)
			}
		})
	},
		func(next http.Handler) http.Handler {
			// Manipulate the header for all the HTTP(S) responses.
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Content-Type", "application/json")

				// Call the next handler, which can be another middleware in the chain, or the final handler.
				next.ServeHTTP(w, r)
			})
		})

	router.PathPrefix("/swagger/").
		Handler(
			httpswagger.Handler(
				httpswagger.DeepLinking(true),
				httpswagger.DocExpansion("none"),
				httpswagger.DomID("#swagger-ui"),
			),
		)
	//port := os.Getenv("PORT")
	port := "8080"

	//rsbuddyClient := rsbuddy.New("https://rsbuddy.com/exchange/summary.json")
	osClient := oldschoolrs.New("https://secure.runescape.com/m=itemdb_oldschool/api/graph")

	items.NewAPI(
		items.Config{
			Router:            router,
			OldschoolRsClient: osClient,
		},
	)

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
		// Timeouts
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
