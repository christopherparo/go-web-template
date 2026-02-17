package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/christopherparo/go-web-template/web"
)

// newRouter builds the application router with all middleware and routes.
func newRouter() http.Handler {
	r := chi.NewRouter()

	// Global middleware stack.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// API routes.
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", handleHealth)
	})

	// Static frontend files.
	staticFS, err := fs.Sub(web.StaticFS, "static")
	if err != nil {
		panic(fmt.Sprintf("failed to create sub filesystem: %v", err))
	}

	fileServer := http.FileServerFS(staticFS)
	r.Handle("/*", fileServer)

	return r
}
