New-Item -ItemType Directory -Force -Path "cmd/api" | Out-Null
New-Item -ItemType File -Force -Path "cmd/api/main.go" | Out-Null

New-Item -ItemType Directory -Force -Path "internal/db" | Out-Null
New-Item -ItemType File -Force -Path "internal/db/db.go" | Out-Null

New-Item -ItemType Directory -Force -Path "internal/handlers" | Out-Null
New-Item -ItemType File -Force -Path "internal/handlers/user_handler_htmx.go" | Out-Null

New-Item -ItemType Directory -Force -Path "internal/models" | Out-Null
New-Item -ItemType File -Force -Path "internal/models/user.go" | Out-Null

New-Item -ItemType Directory -Force -Path "internal/router" | Out-Null
New-Item -ItemType File -Force -Path "internal/router/router.go" | Out-Null

New-Item -ItemType Directory -Force -Path "static" | Out-Null
New-Item -ItemType File -Force -Path "static/index.html" | Out-Null

New-Item -ItemType Directory -Force -Path "static/templates" | Out-Null
New-Item -ItemType File -Force -Path "static/templates/user-list.html" | Out-Null
New-Item -ItemType File -Force -Path "static/templates/user-edit.html" | Out-Null

New-Item -ItemType Directory -Force -Path "pg-init" | Out-Null
New-Item -ItemType File -Force -Path "pg-init/init.sql" | Out-Null

New-Item -ItemType File -Force -Path ".env" | Out-Null
New-Item -ItemType File -Force -Path "Dockerfile" | Out-Null
New-Item -ItemType File -Force -Path "docker-compose.yml" | Out-Null
New-Item -ItemType File -Force -Path "README.md" | Out-Null
