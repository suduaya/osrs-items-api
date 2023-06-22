package items

import (
	"log"
	"net/http"
	"osrs-items-api/internal/items/handler"
	"osrs-items-api/internal/items/service"
	"osrs-items-api/pkg/oldschoolrs"

	"github.com/gorilla/mux"
)

type Config struct {
	Router            *mux.Router
	OldschoolRsClient *oldschoolrs.OldschoolRsClient
}

func NewAPI(c Config) {
	service := service.New(c.OldschoolRsClient)
	handler := handler.New(service)

	setRoutes(handler, c.Router)
}

func setRoutes(h *handler.Handler, router *mux.Router) {
	r := router.PathPrefix("/v1").Subrouter()

	r.HandleFunc("/items/{id}", logTracing(h.GetItemById)).Methods("GET")
}

// Middleware
func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
