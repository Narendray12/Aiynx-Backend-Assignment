# Reasoning – Go Backend Development Task

This document explains my approach, design decisions, and overall implementation for the Go backend development task assigned by Ainyx Solutions.

---

## 1. Problem Understanding

The goal of this assignment was to build a RESTful backend service using Go that:

- Manages users with `name` and `date of birth (dob)`
- Stores `dob` in the database
- Dynamically calculates and returns `age` when fetching user data
- Follows clean backend design practices
- Uses the specified tech stack only (Go, PostgreSQL, SQLC, Fiber, Zap, Validator)

The task also aimed to evaluate adaptability to Go, backend fundamentals, and clarity of thought in design decisions.

---

## 2. High-Level Architecture

I followed a layered (clean) architecture to ensure separation of concerns:

- **Handler Layer** – Handles HTTP requests and responses
- **Service Layer** – Contains business logic (age calculation, orchestration)
- **Repository Layer** – Handles database interactions using SQLC
- **Database Layer** – PostgreSQL (Supabase-compatible)

Request flow:

HTTP Request → Routes → Handler → Service → Repository → Database

This structure improves readability, testability, and long-term maintainability.

---

## 3. Why Age Is Not Stored in the Database

Age is derived data and changes over time. Storing it would lead to:

- Data inconsistency
- Extra update logic (daily/yearly updates)

Instead, I store only `dob` and calculate `age` dynamically using Go’s `time` package. This ensures correctness and keeps the database normalized.

---

## 4. Choice of Tools and Libraries

- **GoFiber**: Lightweight, fast, and well-suited for REST APIs
- **PostgreSQL**: Reliable relational database with strong type support
- **SQLC**: Provides compile-time safety for SQL queries and avoids ORM magic
- **Zap**: Structured logging for better observability
- **go-playground/validator**: Clean input validation at request boundaries

All libraries strictly follow the task requirements.

---

## 5. Database Access with SQLC

SQLC was used to generate type-safe database access code.

Benefits:
- Strong typing between SQL and Go
- Compile-time query validation
- Clear separation of SQL and business logic

Migrations are written explicitly to keep database changes transparent and controlled.

---

## 6. Input Validation Strategy

Validation is handled at the HTTP boundary:

- JSON request body is parsed in handlers
- Fields are validated using `go-playground/validator`
- Date of birth format is validated and parsed (`YYYY-MM-DD`)

Invalid requests fail fast with proper HTTP status codes.

---

## 7. Error Handling & Logging

- Meaningful HTTP status codes are returned (`400`, `404`, `500`, `204`)
- Errors are logged using Zap with context
- Request ID middleware is added to help trace logs per request
- Recovery middleware is used to prevent panics from crashing requests

This makes the application safer and easier to debug.

---

## 8. Environment Configuration

Configuration values such as database URL and port are loaded via environment variables.

- `.env` file is used only for local development
- Production environments inject variables directly
- The application fails fast if required variables are missing

This approach is cloud- and container-friendly.

---

## 9. Dockerization

A multi-stage Docker build is used:

- **Builder stage** compiles the Go binary
- **Runtime stage** runs a minimal Alpine image

Benefits:
- Small image size
- Faster startup
- Secure, production-ready deployment

---

## 10. Supabase Considerations

Supabase is used as a managed PostgreSQL provider.

Key considerations:
- Application tables are created manually (`public.users`)
- Supabase Auth tables (`auth.users`) are untouched
- Row Level Security (RLS) is disabled for development simplicity

---

## 11. Testing Approach

The API was tested using Postman by:

- Verifying all CRUD endpoints
- Testing valid and invalid inputs
- Checking HTTP status codes
- Confirming dynamic age calculation
- Observing logs for each request

---

## 12. Possible Improvements

Given more time, the following enhancements could be added:

- Pagination and filtering for list APIs
- Authentication and authorization
- Swagger/OpenAPI documentation
- Automated unit and integration tests
- CI/CD pipeline setup

---

## 13. Conclusion

This assignment was a great opportunity to apply Go backend concepts in a real-world style setup. The solution focuses on correctness, clarity, and maintainability while strictly adhering to the required tech stack.

Thank you for the opportunity to work on this task.

---