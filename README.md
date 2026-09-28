# Wassimo Catalog Service

Part of the [Wassimo](https://github.com/ahmed-wassim) food-delivery platform.
Owns restaurants, branches, and menus — the current orderable truth.

```text
Repo: github.com/ahmed-wassim/wassimo-catalog
```

## What it owns

Restaurants, branches, menu categories and items, current prices and availability.
Answers: *what can currently be ordered?*

It does not own orders, carts, payments, or deliveries — those are separate
Wassimo services with their own databases.

## Stack

```text
Go 1.27 + Gin v1.12.0 + GORM v1.31.2 + PostgreSQL 18
```

## Run

Requires Docker (PostgreSQL runs in a container).

```powershell
docker compose up -d postgres   # database on localhost:5441
go run ./cmd/api                # API on localhost:8080
```

Or fully in Docker:

```powershell
docker compose up -d --build    # API on localhost:9050
```

## Configuration

```text
PORT            API port (default 8080)
ENV             dev | prod (.env files load only when ENV != prod)
DATABASE_URL    Postgres connection string
```

Copy `.env.example` to `.env` and adjust locally. `.env` is never committed.

## Database

Goose migrations in `database/migrations/`:

```bash
export DATABASE_URL="postgres://catalog:catalog@localhost:5441/catalog_db?sslmode=disable"
goose -dir ./database/migrations postgres "$DATABASE_URL" up
```

## API (Slice 1a)

```text
GET /health                      → {"status":"ok"}
GET /ready                       → {"ready":true|false}
GET /restaurants                 → [{id,name,status}], ?status=active|closed
GET /restaurants/:id/branches    → branches of a restaurant
POST /restaurants/:id/branches   → create a branch (201)
```

Public traffic goes through the Wassimo gateway
(`github.com/ahmed-wassim/wassimo-gateway`), which proxies these routes.

## Layout

```text
cmd/api/             entrypoint: config → database → routes → run
internal/config/     environment-only configuration, fails fast
internal/models/     database mapping (migrations are the source of truth)
internal/repos/      queries: context in, wrapped error out
internal/handlers/   HTTP layer: validation and status codes, no SQL
database/migrations/ versioned SQL schema + seed data
```
