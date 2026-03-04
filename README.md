# Unemployed Days Counter

A simple Go service that counts the number of days since a given date. Built with [Fiber v3](https://github.com/gofiber/fiber).

## Usage

### Embed in your GitHub Profile README

Add this to your `README.md` (replace `YOUR_DOMAIN` and `YOUR_DATE`):

```markdown
![Unemployed Days](https://YOUR_DOMAIN/api/v1/unemploy/badge?start=2025-01-01)
```

Example:

```markdown
![Unemployed Days](https://unemployed.example.com/api/v1/unemploy/svg?start=2025-06-15)
```

### API Endpoints

| Endpoint                                                           | Description                                     |
| ------------------------------------------------------------------ | ----------------------------------------------- |
| `GET /api/v1/unemploy?start=YYYY-MM-DD`                            | Returns an HTML page with the day count         |
| `GET /api/v1/unemploy/badge?start=YYYY-MM-DD`                      | Returns an SVG badge (for embedding)            |
| `GET /api/v1/unemploy/wallpaper?start=YYYY-MM-DD&width=W&height=H` | Returns an SVG wallpaper with custom dimensions |

### Run Locally

```bash
go run cmd/main.go
```

Then visit:

- Card: `http://localhost:8000/api/v1/unemploy?start=2025-01-01`
- Badge: `http://localhost:8000/api/v1/unemploy/badge?start=2025-01-01`
- Wallpaper: `http://localhost:8000/api/v1/unemploy/wallpaper?start=2025-01-01&width=1920&height=1080`

#### Wallpaper Parameters

- `start`: Start date in YYYY-MM-DD format (required)
- `width`: Wallpaper width in pixels (required, min: 800, max: 7680)
- `height`: Wallpaper height in pixels (required, min: 600, max: 4320)

Common resolutions:

- 1920x1080 (Full HD)
- 2560x1440 (2K)
- 3840x2160 (4K)
- 1366x768 (Laptop)

### Run with Docker

```bash
cd docker
docker compose up --build
```

## Tech Stack

- Go 1.25
- Fiber v3
- go-playground/validator v10
