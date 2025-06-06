package router

import (
	"net/http"
	"golangApp-postgres/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter initializes all routes and middleware
func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Static assets served from /static
	fileServer(r, "/static", http.Dir("static"))

	// Redirect root to static/index.html
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/index.html", http.StatusFound)
	})

	// Named routes for each navbar static page
	r.Get("/features", serveStaticPage("features.html"))
	r.Get("/pricing", serveStaticPage("pricing.html"))
	r.Get("/about", serveStaticPage("about.html"))
	r.Get("/contact", serveStaticPage("contact.html"))
	r.Get("/signup", serveStaticPage("signup.html"))
	r.Get("/login", serveStaticPage("login.html"))

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ OK"))
	})

	// HTMX-powered User CRUD
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsersHTMX)
		r.Post("/", handlers.CreateUserHTMX)
		r.Get("/{id}/edit", handlers.EditUserFormHTMX)
		r.Put("/{id}", handlers.UpdateUserHTMX)
		r.Delete("/{id}", handlers.DeleteUserHTMX)
	})

	return r
}

// serveStaticPage returns a handler that serves a static HTML file
func serveStaticPage(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/"+filename)
	}
}

// fileServer serves static files at the given path
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
