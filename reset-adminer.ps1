# Stop containers
Write-Host "🔄 Stopping running containers..." -ForegroundColor Cyan
docker-compose stop
# Remove Adminer to clear its session cache
Write-Host "🧹 Removing Adminer container to reset state..." -ForegroundColor Yellow
docker-compose rm -f adminer

# Optionally, rebuild the app if you made code changes
Write-Host "🔨 Rebuilding the app service..." -ForegroundColor Green
docker-compose build app

# Start all services again
Write-Host "🚀 Starting containers fresh..." -ForegroundColor Cyan
docker-compose up -d

# Show logs from app only
Write-Host "`n📺 Tailing app logs only (use Ctrl+C to stop)..." -ForegroundColor Magenta
docker-compose logs -f app
