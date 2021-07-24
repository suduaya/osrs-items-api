package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"osrs-items-api/oldschoolrs"
	"osrs-items-api/provision"
	"osrs-items-api/rsbuddy"
	"time"

	"github.com/gorilla/mux"
)

type API struct {
	mux    *mux.Router
	server *http.Server

	ItemManager provision.ItemManager
}

// Middleware
func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func New(osrsClient oldschoolrs.OldschoolRsClient, rsbuddyClient rsbuddy.RSBuddyClient) *API {
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

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Timeouts
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	itemManager := provision.ItemManager{
		OldschoolRsClient: osrsClient,
		RSBuddyClient:     rsbuddyClient,
	}

	api := &API{
		mux:    router,
		server: srv,

		ItemManager: itemManager,
	}

	router.HandleFunc("/items/{id}", logTracing(api.GetItemById)).Methods("GET")
	router.HandleFunc("/items", logTracing(api.GetItemByName)).Methods("GET")

	return api
}

func (t *API) Start() error {
	fmt.Println("Starting..")
	return t.server.ListenAndServe()
}

// Shutdown attempts to close the http server.
func (t *API) Close() error {
	return t.server.Shutdown(context.Background())
}
