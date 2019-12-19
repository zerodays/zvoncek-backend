package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"zvon/config"
	"zvon/logger"
)

// Middleware for logging server activity.
func loggerMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		next(w, r, ps)

		logger.Log(fmt.Sprintf("%s\t%s\t%dms", r.Method, r.RequestURI, time.Since(start).Milliseconds()), "ROUTER", logger.LevelInfo)
	}
}

// Helper function that writes CORS header.
func writeCorsHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization")
}

// Handle CORS (OPTIONS) requests. Writes CORS header fields
// and responds with empty body and status 200.
func corsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writeCorsHeader(w)

	_, _ = w.Write([]byte{})
}

// Middleware that writes CORS headers before further processing the request.
func corsMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		writeCorsHeader(w)

		next(w, r, ps)
	}
}

// Middleware that writes Content-Type header before further processing the request.
func responseTypeHeaderMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		next(w, r, ps)
	}
}

func addMiddleware(handler httprouter.Handle, customContentType bool) httprouter.Handle {
	// NOTE: Middleware is executed from bottom to top!!

	handler = corsMiddleware(handler)

	if !customContentType {
		handler = responseTypeHeaderMiddleware(handler)
	}

	if config.Current.MinimumLogLevel <= int(logger.LevelInfo) {
		handler = loggerMiddleware(handler)
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handler(w, r, ps)
	}
}

func addRoute(router *httprouter.Router, prefix string, r Route) {
	path := fmt.Sprintf("%s%s", prefix, r.Path)

	// Add handlers if needed
	if r.GET != nil {
		router.Handle("GET", path, addMiddleware(r.GET, r.CustomContentType))
	}

	if r.PUT != nil {
		router.Handle("PUT", path, addMiddleware(r.PUT, r.CustomContentType))
	}

	if r.POST != nil {
		router.Handle("POST", path, addMiddleware(r.POST, r.CustomContentType))
	}

	if r.DELETE != nil {
		router.Handle("DELETE", path, addMiddleware(r.DELETE, r.CustomContentType))
	}

	// Add CORS handler for OPTIONS request
	router.Handle("OPTIONS", path, addMiddleware(corsHandler, false))
}

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	for _, r := range routes {
		addRoute(router, "/api/v1", r)
	}

	return router
}
