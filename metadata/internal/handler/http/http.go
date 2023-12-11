package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/thernande/movie-micro/metadata/internal/controller/metadata"
	"github.com/thernande/movie-micro/metadata/internal/repository"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	ctrl *metadata.Controller
}

// New creates a new movie metadata HTTP handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl}
}

// GetMetadata handles GET /metadata requests.
func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// `ctx := req.Context()` is creating a new context for the request. The context is used to carry
	// request-scoped values across API boundaries and between middleware and handlers. It allows passing
	// values such as request-specific data, cancellation signals, and deadlines to functions and
	// goroutines involved in handling the request.
	ctx := req.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
