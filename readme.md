# AIYNX BACKEND ASSIGNMENT – User Management API

User Management API is a production-ready RESTful API built with **Go** and **Fiber** for managing users with their **name** and **date of birth (DOB)**.  
The API dynamically calculates a user’s **age** at runtime instead of storing it in the database.

---

## Features

- RESTful CRUD APIs for users
- Dynamic age calculation from DOB
- PostgreSQL database (Supabase compatible)
- SQLC for type-safe SQL queries
- Input validation using `go-playground/validator`
- Structured logging with Uber Zap
- Request ID and request duration middleware
- Environment-based configuration using `.env`
- Production-ready Docker setup

---

## Tech Stack

| Category | Technology |
|--------|------------|
| Language | Go (1.25) |
| Framework | GoFiber |
| Database | PostgreSQL (Supabase) |
| SQL Layer | SQLC |
| Validation | go-playground/validator |
| Logging | Uber Zap |
| Containerization | Docker |

---

## Project Structure

| Path | Description |
|-----|------------|
| `cmd/server/main.go` | Application entry point |
| `config/` | Environment configuration loader |
| `db/` | Database related files |
| `db/migrations/` | SQL database migrations |
| `db/sqlc/` | SQLC schema and generated code |
| `internal/` | Internal application logic |
| `internal/handler/` | HTTP request handlers |
| `internal/service/` | Business logic layer |
| `internal/repository/` | Database access layer |
| `internal/routes/` | API route definitions |
| `internal/middleware/` | Custom middleware (request ID, logging, recovery) |
| `internal/models/` | Domain models |
| `internal/logger/` | Logger configuration (Zap) |
| `.env` | Environment variables (local only, not committed) |
| `Dockerfile` | Docker configuration for production build |
| `go.mod` | Go module definition |
| `go.sum` | Go module dependency checksums |

---

## Database Schema

```sql
CREATE TABLE public.users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

**Note**
- This table must be created manually in Supabase using the SQL Editor.
- Do **not** modify `auth.users` (Supabase internal table).

---

## API Endpoints

| Method | Endpoint | Description |
|------|---------|-------------|
| POST | `/users` | Create a new user |
| GET | `/users` | List all users |
| GET | `/users/{id}` | Get user by ID |
| PUT | `/users/{id}` | Update user |
| DELETE | `/users/{id}` | Delete user |

---

## Environment Variables

| Variable | Description |
|--------|------------|
| `APP_ENV` | Application environment |
| `APP_PORT` | Port number |
| `DATABASE_URL` | PostgreSQL connection string |

```env
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgresql://postgres:<password>@db.<project>
```

---

## Running Locally

```bash
go mod tidy
go run ./cmd/server
```

Server runs at:  
`http://localhost:8080`

---

## Running with Docker

```bash
docker build -t aiynx-api .
docker run -p 8080:8080 --env-file .env aiynx-api
```
