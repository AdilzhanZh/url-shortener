# URL Shortener

A small Go URL shortener API built with Gin and PostgreSQL.

## Features

- Create short links for `http` and `https` URLs
- Optionally provide a custom short code
- Redirect short links with `302 Found`
- Track click count in PostgreSQL
- Expire links 7 days after creation
- Health check endpoint
- Docker Compose setup with database migrations

## Tech Stack

- Go 1.25
- Gin
- PostgreSQL
- sqlx with pgx
- golang-migrate
- Docker and Docker Compose

## Project Structure

```text
cmd/api/                 Application entrypoint
internal/config/         Environment configuration
internal/handler/        HTTP handlers and routes
internal/model/          Data models
internal/repository/     PostgreSQL access layer
internal/server/         HTTP server wrapper
internal/service/        URL business logic
migrations/              Database migrations
```

## Configuration

Create a `.env` file from the example:

```powershell
Copy-Item .env.example .env
```

Environment variables:

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | API server port |
| `LOG_LEVEL` | `INFO` | Logger level |
| `DB_HOST` | `localhost` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `postgres` | PostgreSQL user |
| `DB_PASSWORD` | `postgres` | PostgreSQL password |
| `DB_NAME` | `postgres` | PostgreSQL database name |
| `SSL_MODE` | `disable` | PostgreSQL SSL mode |

For Docker Compose, `DB_HOST` is set to `postgres-db` inside `docker-compose.yml`.

## Run With Docker Compose

Start the API and PostgreSQL:

```powershell
docker compose up --build
```

The backend container runs database migrations before starting the API.

API URL:

```text
http://localhost:8080
```

Stop the containers:

```powershell
docker compose down
```

To also remove the PostgreSQL volume:

```powershell
docker compose down -v
```

## Run Locally

Start PostgreSQL separately, then set `.env` values for your local database. For local execution, `DB_HOST` should usually be:

```env
DB_HOST=localhost
```

Run migrations:

```powershell
migrate -path migrations -database "postgres://admin:admin@localhost:5432/url_shortener_db?sslmode=disable" up
```

Start the API:

```powershell
go run ./cmd/api
```

## API

### Health Check

```http
GET /health
```

Response:

```json
{
  "status": "ok"
}
```

### Create Short URL

```http
POST /api/shorten
Content-Type: application/json
```

Request with generated short code:

```json
{
  "original_url": "https://example.com"
}
```

Request with custom short code:

```json
{
  "original_url": "https://example.com",
  "short_code": "example"
}
```

Response:

```json
{
  "short_url": "http://localhost:8080/example"
}
```

### Redirect

```http
GET /:code
```

Example:

```http
GET /example
```

Redirects to the original URL with `302 Found`.

## Examples

Create a short URL:

```powershell
curl.exe -X POST http://localhost:8080/api/shorten `
  -H "Content-Type: application/json" `
  -d "{\"original_url\":\"https://example.com\",\"short_code\":\"example\"}"
```

Open the shortened URL:

```powershell
curl.exe -i http://localhost:8080/example
```

## Database

The `urls` table stores:

- Original URL
- Unique short code
- Creation and update timestamps
- Expiration timestamp
- Click count

New links are assigned an expiration date of 7 days after creation.
