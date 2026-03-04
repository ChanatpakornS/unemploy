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

| Endpoint                                                           | Description                                    |
| ------------------------------------------------------------------ | ---------------------------------------------- |
| `GET /api/v1/unemploy?start=YYYY-MM-DD`                            | Returns an SVG card with animated day count    |
| `GET /api/v1/unemploy/badge?start=YYYY-MM-DD`                      | Returns an SVG badge (for embedding)           |
| `GET /api/v1/unemploy/wallpaper?start=YYYY-MM-DD&width=W&height=H` | Returns a PNG wallpaper with custom dimensions |

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
- Fiber v3 (web framework)
- go-playground/validator v10 (request validation)
- gg (2D graphics library for direct PNG rendering)
- freetype (font rendering)
- Air (live reloading for development)

## Features

### 🎴 Card View

- Animated SVG counter with smooth transitions
- Clean, modern design (256×192)
- Perfect for GitHub profile READMEs

### 🏷️ Badge View

- Compact SVG badge format
- Auto-sizing based on day count
- Gradient styling with animations

### 🖼️ Wallpaper Generator

- **PNG format** - Ready to use as desktop/mobile wallpaper (directly rendered, not converted from SVG)
- Custom resolutions (800-7680 × 600-4320 pixels)
- Responsive font sizing that adapts to your screen
- Beautiful gradient background (slate colors)
- Large, easy-to-read day counter
- Glowing text effects for visual appeal
- Clean, modern typography

> **Note:** The wallpaper uses direct PNG rendering for crisp, clear text at any resolution. The day count is displayed statically (no animation in PNG format).

## Download Wallpaper Examples

```bash
# Full HD (1920×1080)
curl "http://localhost:8000/api/v1/unemploy/wallpaper?start=2025-01-01&width=1920&height=1080" -o wallpaper.png

# 4K (3840×2160)
curl "http://localhost:8000/api/v1/unemploy/wallpaper?start=2025-01-01&width=3840&height=2160" -o wallpaper-4k.png

# Mobile Portrait (1080×1920)
curl "http://localhost:8000/api/v1/unemploy/wallpaper?start=2025-01-01&width=1080&height=1920" -o wallpaper-mobile.png
```
