# 🚀 Go Web MVP Template

A clean, functional starter template for building minimal web applications using:

- ⚙️ **Go** for backend logic
- ⚡ **HTMX** for modern interactions without JavaScript frameworks
- 🐘 **PostgreSQL** for persistent data
- 🎨 **Bulma CSS** for a lightweight UI
- 🖥️ **Adminer** for simple database management

---

## 📁 File Structure

```

├── cmd/
│   └── api/
│       └── main.go              # Entry point of the Go web server
├── internal/
│   ├── db/
│   │   └── db.go                # DB connection + queries
│   ├── handlers/
│   │   └── user\_handler\_htmx.go # HTMX-compatible CRUD handlers
│   ├── models/
│   │   └── user.go              # User struct definition
│   └── router/
│       └── router.go            # Chi router setup
├── static/
│   ├── index.html               # Frontend entry page (HTMX + Bulma)
│   └── templates/
│       ├── user-list.html       # User list fragment (HTMX)
│       └── user-edit.html       # Edit form fragment (HTMX)
├── pg-init/
│   └── init.sql                 # Bootstrap SQL for PostgreSQL
├── .env                         # Environment variables for DB config
├── .gitignore                   # Git ignore rules
├── Dockerfile                   # Builds the Go app container
├── docker-compose.yml           # Full development stack orchestration
├── reset-adminer.ps1            # Utility script to reset Adminer
├── New-GoHTMX-Postgres-Template.ps1 # PowerShell script to scaffold a new app
└── README.md                    # This file

````

---

## 🧪 Running the App (Local Dev)

```powershell
# 1. Ensure Docker is running
# 2. From the project root:
docker-compose up --build
````

* App runs at: [http://localhost:8080](http://localhost:8080)
* Adminer GUI: [http://localhost:8081](http://localhost:8081)

Login in Adminer with:

* System: PostgreSQL
* Server: postgres
* Username: lesson\_pg\_user
* Password: lesson\_pg\_pass
* Database: lesson\_pg\_db

---

## 🌱 Starting a New Project from This Template

```powershell
# Clone the base
Copy-Item -Recurse -Force .\go-web-mvp-template .\your-new-project

# Enter new folder
Set-Location .\your-new-project

# Remove Git history and init new repo
Remove-Item -Recurse -Force .git
git init
git remote add origin https://github.com/your-username/your-new-project.git

# Push your initial commit
git add .
git commit -m "🌱 Initial commit from go-web-mvp-template"
git branch -M main
git push -u origin main
```

---

## 🛠️ Customization Tips

* To add a new model:

  * Create a struct in `internal/models/`
  * Add DB functions in `internal/db/`
  * Add route handlers in `internal/handlers/`
  * Update `router/router.go`
  * Create `.html` templates in `static/templates`

* Change the frontend styles in `static/index.html` using Bulma.

* Replace HTMX fragments with your custom HTML and Go handlers.

---

## 🔐 Security Notice

* `.env` is ignored via `.gitignore` and **must never** be committed.
* Do **not** use the default credentials in production.
* For production, enforce HTTPS, validate all inputs, and restrict DB access.

---

## 📄 License

MIT – Use freely for personal and commercial projects.

```