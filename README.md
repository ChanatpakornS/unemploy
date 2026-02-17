# Unemployed Days Counter

A simple Go service that counts the number of days since a given date. Built with [Fiber v3](https://github.com/gofiber/fiber).

## Usage

### Embed in your GitHub Profile README

Add this to your `README.md` (replace `YOUR_DOMAIN` and `YOUR_DATE`):

```markdown
![Unemployed Days](https://YOUR_DOMAIN/api/v1/unemploy/svg?start=2025-01-01)
```

Example:

```markdown
![Unemployed Days](https://unemployed.example.com/api/v1/unemploy/svg?start=2025-06-15)
```

### API Endpoints

| Endpoint                                    | Description                             |
| ------------------------------------------- | --------------------------------------- |
| `GET /api/v1/unemploy?start=YYYY-MM-DD`     | Returns an HTML page with the day count |
| `GET /api/v1/unemploy/svg?start=YYYY-MM-DD` | Returns an SVG badge (for embedding)    |

### Run Locally

```bash
go run cmd/main.go
```

Then visit: `http://localhost:8000/api/v1/unemploy?start=2025-01-01`

### Run with Docker

```bash
cd docker
docker compose up --build
```

## Tech Stack

- Go 1.25
- Fiber v3
- go-playground/validator v10
