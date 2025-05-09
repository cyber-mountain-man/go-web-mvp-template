package handlers

import (
	"net/http"
	"strconv"
	"html/template"
	"log"

	"golangApp-postgres/internal/db"

	"github.com/go-chi/chi/v5"
)

var templates = template.Must(template.ParseGlob("static/templates/*.html"))

// ListUsersHTMX handles GET /users
func ListUsersHTMX(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "user-list.html", users)
	if err != nil {
		log.Printf("Template execution failed: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

// CreateUserHTMX handles POST /users
func CreateUserHTMX(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form submission", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := db.InsertUser(name, email); err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	ListUsersHTMX(w, r)
}

// EditUserFormHTMX handles GET /users/{id}/edit
func EditUserFormHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	err = templates.ExecuteTemplate(w, "user-edit.html", user)
	if err != nil {
		log.Printf("Edit template error: %v", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

// UpdateUserHTMX handles PUT /users/{id}
func UpdateUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := db.UpdateUser(id, name, email); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Notify frontend to clear the edit form
	w.Header().Set("HX-Trigger", "userUpdated")

	// Refresh user List
	ListUsersHTMX(w, r)
}

// DeleteUserHTMX handles DELETE /users/{id}
func DeleteUserHTMX(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := db.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	ListUsersHTMX(w, r)
}
