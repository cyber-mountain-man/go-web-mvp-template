
# 🚀 HTMX + Go + PostgreSQL Starter Template

This project is a full-stack web application starter built with:

- ✅ Go (chi router)
- ✅ PostgreSQL for persistent storage
- ✅ HTMX for seamless frontend interactions
- ✅ Docker & Docker Compose for local development
- ✅ Adminer for PostgreSQL GUI management

---

## 📁 Project Structure

```
├── cmd/
│   └── api/
│       └── main.go              # App entry point
├── internal/
│   ├── db/
│   │   └── db.go                # DB connection and queries
│   ├── handlers/
│   │   └── user_handler_htmx.go # HTMX route handlers
│   ├── models/
│   │   └── user.go              # User model
│   └── router/
│       └── router.go            # Chi router configuration
├── static/
│   ├── index.html               # Main frontend page
│   └── templates/
│       ├── user-list.html       # List fragment
│       └── user-edit.html       # Edit form
├── pg-init/
│   └── init.sql                 # Initializes DB schema/data
├── .env                         # Environment variables
├── Dockerfile                   # Go build configuration
├── docker-compose.yml           # Development environment config
└── README.md                    # This file
```

---

## ⚙️ Getting Started

### 🧰 Prerequisites

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/cyber-mountain-man/golangApp-postgres.git
cd golangApp-postgres
```

### 2. Create `.env` File

```env
DBUSER=lesson_pg_user
DBPASSWORD=lesson_pg_pass
DBHOST=postgres
DBPORT=5432
DBNAME=lesson_pg_db
```

### 3. Run with Docker Compose

```bash
docker compose up --build
```

Your services should be live:

- HTMX Frontend: http://localhost:8080
- Adminer DB GUI: http://localhost:8081

---

## 🛠️ Adminer Login

- System: PostgreSQL
- Server: postgres
- Username: `lesson_pg_user`
- Password: `lesson_pg_pass`
- Database: `lesson_pg_db`

---

## 🐘 Verify DB with PSQL (optional)

```powershell
docker run -it --rm --network=golangapp-postgres_go-net postgres:15   psql -h postgres -U lesson_pg_user -d lesson_pg_db
```

Once inside:

```sql
SELECT * FROM users;
```

---

## 🧼 Tear Down

```bash
docker compose down -v
```

---

## 📌 Notes

- This template uses multi-stage Docker builds for minimal image size.
- Adminer is included for convenience in development.
- You can use this structure to scaffold future apps by replacing `users` with your own models.

---

## ⚠️ Security Notice

This project is for **local development only**. Please read carefully:

### 1. Never commit `.env` files to Git
They contain sensitive DB credentials.

### 2. Hardcoded Credentials Are Public
Replace them in production and manage securely.

### 3. Adminer is for development
Never expose it on public deployments.

### 4. This app has no authentication
Anyone with access can modify your data.

### 5. Keep PostgreSQL ports private
Don’t expose port 5432 externally without firewall/proxy protection.

---

 **In short**: Keep this app on your local machine or secure everything before deploying.
