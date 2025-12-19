# 🚀 AIYNX Backend Assignment - User Management API

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=fiber)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

A production-ready RESTful API for user management built with Go and Fiber, featuring dynamic age calculation, clean architecture, and full containerization.

[Features](#-features) • [Quick Start](#-quick-start) • [API Documentation](#-api-endpoints) • [Architecture](#-project-structure)

</div>

---

## ✨ Features

- 🔄 **Full CRUD Operations** - Complete user management functionality
- 🎂 **Dynamic Age Calculation** - Real-time age computation from date of birth
- 🗄️ **PostgreSQL Database** - Robust and scalable data storage
- 🔒 **Type-Safe Queries** - SQLC for compile-time SQL validation
- ✅ **Input Validation** - go-playground/validator for request validation
- 📊 **Structured Logging** - Uber Zap for production-grade logging
- 🎯 **Request Tracking** - Middleware for request IDs and duration logging
- 🌍 **Environment Configuration** - Flexible .env-based configuration
- 🐳 **Fully Dockerized** - Ready for containerized deployment

---

## 🛠️ Tech Stack

| Category | Technology |
|----------|-----------|
| **Language** | Go 1.25 |
| **Framework** | GoFiber |
| **Database** | PostgreSQL (Supabase) |
| **SQL Layer** | SQLC |
| **Validation** | go-playground/validator |
| **Logging** | Uber Zap |
| **Containerization** | Docker |

---

## 📁 Project Structure
```
AIYNX/
├── 📂 cmd/
│   └── 📂 server/
│       └── 📄 main.go              # Application entry point
├── 📂 config/                      # Environment configuration loader
├── 📂 db/
│   ├── 📂 migrations/              # Database migrations
│   └── 📂 sqlc/                    # SQLC schema and generated code
├── 📂 internal/
│   ├── 📂 handler/                 # HTTP handlers
│   ├── 📂 service/                 # Business logic layer
│   ├── 📂 repository/              # Database access layer
│   ├── 📂 routes/                  # Route definitions
│   ├── 📂 middleware/              # Custom middleware
│   ├── 📂 models/                  # Domain models
│   └── 📂 logger/                  # Logger configuration
├── 📄 .env                         # Environment variables (local only)
├── 📄 Dockerfile                   # Docker configuration
├── 📄 go.mod
└── 📄 go.sum
```

---

## 🗄️ Database Schema
```sql
CREATE TABLE public.users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

---

## 🌐 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/users` | Create a new user |
| `GET` | `/users` | Retrieve all users |
| `GET` | `/users/{id}` | Retrieve a specific user by ID |
| `PUT` | `/users/{id}` | Update an existing user |
| `DELETE` | `/users/{id}` | Delete a user |

### 📝 Request/Response Examples

<details>
<summary><b>POST /users - Create User</b></summary>

**Request:**
```json
{
  "name": "John Doe",
  "dob": "1990-05-15"
}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-05-15",
  "age": 34,
  "created_at": "2025-12-19T10:30:00Z",
  "updated_at": "2025-12-19T10:30:00Z"
}
```
</details>

<details>
<summary><b>GET /users/{id} - Get User</b></summary>

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-05-15",
  "age": 34,
  "created_at": "2025-12-19T10:30:00Z",
  "updated_at": "2025-12-19T10:30:00Z"
}
```
</details>

---

## 🔧 Environment Variables

Create a `.env` file in the root directory:
```env
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgresql://postgres:<password>@db.<project>.supabase.co:5432/postgres
```

> ⚠️ **Note:** The `.env` file is not committed to version control. In production, inject environment variables directly.

---

## 🚀 Quick Start

### Prerequisites

- Go 1.25 or higher
- PostgreSQL database
- Docker (optional, for containerized deployment)

### Local Development

1. **Clone the repository**
```bash
   git clone https://github.com/yourusername/aiynx-backend.git
   cd aiynx-backend
```

2. **Install dependencies**
```bash
   go mod tidy
```

3. **Set up environment variables**
```bash
   cp .env.example .env
   # Edit .env with your configuration
```

4. **Run the application**
```bash
   go run ./cmd/server
```

5. **Access the API**
```
   http://localhost:8080
```

### 🐳 Docker Deployment

1. **Build the Docker image**
```bash
   docker build -t aiynx-api .
```

2. **Run the container**
```bash
   docker run -p 8080:8080 --env-file .env aiynx-api
```

3. **Access the API**
```
   http://localhost:8080
```

---

## 🏗️ Architecture & Design Decisions

### Clean Architecture Principles
```
┌─────────────────────────────────────┐
│         HTTP Layer (Fiber)          │
│            (Handlers)               │
├─────────────────────────────────────┤
│        Business Logic Layer         │
│            (Services)               │
├─────────────────────────────────────┤
│       Data Access Layer             │
│         (Repository)                │
├─────────────────────────────────────┤
│      Database Layer (SQLC)          │
│          (PostgreSQL)               │
└─────────────────────────────────────┘
```

### Key Design Choices

-  **Dynamic Age Calculation** - Age is computed at runtime to prevent stale data
-  **Separation of Concerns** - Business logic isolated in service layer
-  **Type Safety** - SQLC ensures compile-time SQL validation
-  **No ORM Magic** - Explicit SQL for better performance and control
-  **Fail-Fast Configuration** - Early validation of environment setup
-  **Observability First** - Structured logging for production monitoring
-  **Testable Code** - Clean architecture enables easy unit testing

---

## 📊 Logging & Monitoring

The API includes comprehensive logging:

- ✅ Request IDs for distributed tracing
- ⏱️ Request duration tracking
- 📝 Structured JSON logs via Uber Zap
- 🎯 Error tracking with context

---