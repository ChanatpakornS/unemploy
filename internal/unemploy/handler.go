package unemploy

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/unemploy", Unemploy)
	v1.Get("/unemploy/svg", UnemploySVG)
}

func Unemploy(c fiber.Ctx) error {
	dateRequest := new(UnemployedRequestParams)
	if err := c.Bind().Query(dateRequest); err != nil {
		log.Error("Failed to get date value", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	start, err := time.Parse("2006-01-02", dateRequest.Start)
	if err != nil {
		log.Error("Failed to parse date", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	now := time.Now()
	diff := now.Sub(start)
	if diff < 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Date cannot be in the future",
		})
	}

	days := int(diff.Hours() / 24)

	return c.Render("index", fiber.Map{
		"Days": days,
	})
}

func UnemploySVG(c fiber.Ctx) error {
	dateRequest := new(UnemployedRequestParams)
	if err := c.Bind().Query(dateRequest); err != nil {
		log.Error("Failed to get date value", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	start, err := time.Parse("2006-01-02", dateRequest.Start)
	if err != nil {
		log.Error("Failed to parse date", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	now := time.Now()
	diff := now.Sub(start)
	if diff < 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Date cannot be in the future",
		})
	}

	days := int(diff.Hours() / 24)
	daysStr := fmt.Sprintf("%d", days)

	// Calculate dynamic width based on digit count
	valueText := fmt.Sprintf("%s days", daysStr)
	labelWidth := 96
	valueWidth := 20 + len(valueText)*7
	totalWidth := labelWidth + valueWidth

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="28" role="img" aria-label="unemployed: %s">
  <defs>
    <linearGradient id="bg" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0%%" stop-color="#2b2d42"/>
      <stop offset="100%%" stop-color="#1a1a2e"/>
    </linearGradient>
    <linearGradient id="val" x1="0" y1="0" x2="1" y2="0">
      <stop offset="0%%" stop-color="#f7971e"/>
      <stop offset="100%%" stop-color="#ffd200"/>
    </linearGradient>
    <filter id="shadow">
      <feDropShadow dx="0" dy="1" stdDeviation="1" flood-opacity="0.25"/>
    </filter>
  </defs>
  <rect width="%d" height="28" rx="6" fill="url(#bg)" filter="url(#shadow)"/>
  <rect x="%d" width="%d" height="28" rx="6" fill="url(#val)"/>
  <rect x="%d" width="6" height="28" fill="url(#val)"/>
  <g font-family="'Segoe UI',Roboto,Verdana,sans-serif" font-size="11" font-weight="600">
    <text x="%d" y="18" fill="#ffffff" text-anchor="middle">unemployed</text>
    <text x="%d" y="18" fill="#000009" text-anchor="middle" font-weight="800">%s</text>
  </g>
</svg>`,
		totalWidth, valueText,
		totalWidth,
		labelWidth, valueWidth,
		labelWidth,
		labelWidth/2,
		labelWidth+valueWidth/2, valueText,
	)

	c.Set("Content-Type", "image/svg+xml")
	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	return c.SendString(svg)
}
