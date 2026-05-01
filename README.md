# Todo App - Full-Stack Todo Management Application

A React + Vite + Tailwind 4 frontend with a Go + chi backend for managing todos with user authentication.

## Prerequisites

### Backend
- Go 1.23 or higher

### Frontend
- Node.js 18+ with npm

## Project Structure

```
.
 backend/           # Go REST API server
   ├── cmd/server/   # Server entrypoint
   ├── migrations/    # Database migration files (golang-migrate)
   └── go.mod         # Go module definition
 frontend/          # React + Vite + Tailwind 4 application
    ├── src/           # React components and pages
    ├── package.json   # NPM dependencies
    └── vite.config.js # Vite configuration
```

## Running the Services

### Backend

```bash
cd backend
go run ./cmd/server
```

The backend will start on `http://localhost:8080`

### Frontend

```bash
cd frontend
npm install  # Only needed on first run or after dependency changes
npm run dev
```

The frontend will start on `http://localhost:5173`

## Architecture

- **Backend**: Go chi router with PostgreSQL support (via sqlx and golang-migrate)
  - Health endpoint: `GET /api/health` → `{"status":"ok"}`
  - CORS configured for `http://localhost:5173`

- **Frontend**: React application with Tailwind 4 CSS-first styling
  - Vite for fast development and builds
  - Fetches health status from backend on load
  - Environment-driven API URL configuration

## API Endpoints

### Health Check
- `GET /api/health` - Backend health status
  - Response: `{"status":"ok"}`

## Development Workflow

1. Start the backend: `cd backend && go run ./cmd/server`
2. In a new terminal, start the frontend: `cd frontend && npm run dev`
3. Open http://localhost:5173 in your browser
4. The frontend will display the backend health status once the backend is running

## CORS Configuration

The backend allows requests from `http://localhost:5173` during development. Modify the CORS configuration in `backend/main.go` if needed.
