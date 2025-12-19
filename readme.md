# AIYNX BACKEND ASSIGNMENT -- User Management API

User Management API is a production-ready RESTful API built using Go and Fiber for
managing users with name and date of birth (DOB). The API dynamically
calculates user age at runtime instead of storing it in the database. It
follows clean architecture principles and is designed to be scalable,
maintainable, and deployment-ready.

------------------------------------------------------------------------

## Features

-   CRUD APIs for user management
-   Dynamic age calculation using DOB
-   PostgreSQL database 
-   SQLC for type-safe database queries
-   Input validation with go-playground/validator
-   Structured logging using Uber Zap
-   Middleware for request ID and request duration logging
-   Environment-based configuration using .env
-   Fully Dockerized setup

------------------------------------------------------------------------

## Tech Stack

-   Language: Go (1.25)
-   Framework: GoFiber
-   Database: PostgreSQL (Supabase)
-   SQL Layer: SQLC
-   Validation: go-playground/validator
-   Logging: Uber Zap
-   Containerization: Docker

------------------------------------------------------------------------

## Project Structure

AIYNX/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── config/                  # Environment configuration loader
├── db/
│   ├── migrations/          # Database migrations
│   └── sqlc/                # SQLC schema and generated code
├── internal/
│   ├── handler/             # HTTP handlers
│   ├── service/             # Business logic
│   ├── repository/          # Database access layer
│   ├── routes/              # Route definitions
│   ├── middleware/          # Custom middleware
│   ├── models/              # Domain models
│   └── logger/              # Logger configuration
├── .env                     # Environment variables (local only)
├── Dockerfile               # Docker configuration
├── go.mod
└── go.sum

------------------------------------------------------------------------

## Database Schema

CREATE TABLE public.users ( id BIGSERIAL PRIMARY KEY, name TEXT NOT
NULL, dob DATE NOT NULL, created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ NOT NULL DEFAULT now() );

------------------------------------------------------------------------

## API Endpoints

POST /users\
Creates a new user

GET /users\
Returns all users

GET /users/{id}\
Returns a user by ID

PUT /users/{id}\
Updates a user

DELETE /users/{id}\
Deletes a user

------------------------------------------------------------------------

## Environment Variables

APP_ENV=development\
APP_PORT=8080\
DATABASE_URL=postgresql://postgres:`<password>`{=html}@db.`<project>`

.env file is not committed. In production, environment variables should
be injected directly.

------------------------------------------------------------------------

## Running Locally

go mod tidy\
go run ./cmd/server

Server runs on http://localhost:8080

------------------------------------------------------------------------

## Running with Docker

docker build -t aiynx-api .\
docker run -p 8080:8080 --env-file .env aiynx-api

------------------------------------------------------------------------

## Design Decisions

-   Age is calculated dynamically to avoid stale data
-   Business logic lives in the service layer
-   Repository layer uses SQLC for compile-time safety
-   Explicit SQL instead of ORM magic
-   Fail-fast configuration loading
-   Structured logs for observability