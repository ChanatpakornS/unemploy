package unemploy

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/unemploy", Unemploy)
	v1.Get("/unemploy/badge", UnemployBadge)
}

func parseAndCalculateDays(c fiber.Ctx) (int, error) {
	dateRequest := new(UnemployedRequestParams)
	if err := c.Bind().Query(dateRequest); err != nil {
		log.Error("Failed to get date value", err)
		return 0, c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	start, err := time.Parse("2006-01-02", dateRequest.Start)
	if err != nil {
		log.Error("Failed to parse date", err)
		return 0, c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid date format",
		})
	}

	now := time.Now()
	diff := now.Sub(start)
	if diff < 0 {
		return 0, c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Date cannot be in the future",
		})
	}

	return int(diff.Hours() / 24), nil
}

func generateCountingFrames(days int) (cssStr string, textsStr string) {
	totalFrames := 60
	totalDuration := 2.0
	frameDuration := totalDuration / float64(totalFrames)

	var frames []int
	seen := map[int]bool{}
	for i := 1; i <= totalFrames; i++ {
		t := float64(i) / float64(totalFrames)
		eased := 1.0 - math.Pow(2, -10*t)
		val := int(math.Round(eased * float64(days)))
		if seen[val] && i != totalFrames {
			continue
		}
		seen[val] = true
		frames = append(frames, val)
	}
	frames[len(frames)-1] = days

	actualFrames := len(frames)
	var css, texts strings.Builder
	for i, val := range frames {
		delay := float64(i) * frameDuration
		if i == actualFrames-1 {
			css.WriteString(fmt.Sprintf(
				"    .f%d { animation: show %.3fs ease forwards; animation-delay: %.3fs; opacity: 0; }\n",
				i, frameDuration, delay))
		} else {
			css.WriteString(fmt.Sprintf(
				"    .f%d { animation: hide %.3fs ease forwards; animation-delay: %.3fs; opacity: 0; }\n",
				i, frameDuration, delay))
		}
		texts.WriteString(fmt.Sprintf(
			`      <text class="f%d" x="128" y="112" text-anchor="middle" dominant-baseline="central" font-size="56" font-weight="800" fill="#1a73e8">%d</text>`+"\n",
			i, val))
	}
	return css.String(), texts.String()
}

func Unemploy(c fiber.Ctx) error {
	days, err := parseAndCalculateDays(c)
	if err != nil {
		return err
	}

	cssFrames, textFrames := generateCountingFrames(days)

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="256" height="192" role="img" aria-label="unemployed for %d days">
  <style>
    @keyframes show { from { opacity: 0; } to { opacity: 1; } }
    @keyframes hide { 0%% { opacity: 0; } 10%% { opacity: 1; } 90%% { opacity: 1; } 100%% { opacity: 0; } }
%s  </style>
  <rect width="256" height="192" rx="16" fill="#ffffff"/>
  <rect x="1" y="1" width="254" height="190" rx="15" fill="none" stroke="#d2e3fc" stroke-width="2"/>
  <g font-family="'Segoe UI',Roboto,Verdana,sans-serif">
    <text x="128" y="52" text-anchor="middle" font-size="11" fill="#5f6f81" letter-spacing="1.5" text-transform="uppercase">I'VE BEEN UNEMPLOYED FOR</text>
%s    <text x="128" y="158" text-anchor="middle" font-size="11" fill="#5f6f81" letter-spacing="1.5">DAYS</text>
  </g>
</svg>`,
		days,
		cssFrames,
		textFrames,
	)

	c.Set("Content-Type", "image/svg+xml")
	c.Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=86400")
	return c.SendString(svg)
}

func UnemployBadge(c fiber.Ctx) error {
	days, err := parseAndCalculateDays(c)
	if err != nil {
		return err
	}

	daysStr := fmt.Sprintf("%d", days)

	totalFrames := 60
	totalDuration := 2.0
	frameDuration := totalDuration / float64(totalFrames)

	var frames []int
	seen := map[int]bool{}
	for i := 1; i <= totalFrames; i++ {
		t := float64(i) / float64(totalFrames)
		eased := 1.0 - math.Pow(2, -10*t)
		val := int(math.Round(eased * float64(days)))
		if seen[val] && i != totalFrames {
			continue
		}
		seen[val] = true
		frames = append(frames, val)
	}
	frames[len(frames)-1] = days

	valueText := fmt.Sprintf("%s days", daysStr)
	labelWidth := 96
	valueWidth := 20 + len(valueText)*7
	totalWidth := labelWidth + valueWidth
	valX := labelWidth + valueWidth/2

	actualFrames := len(frames)
	var css, texts strings.Builder
	for i, val := range frames {
		delay := float64(i) * frameDuration
		if i == actualFrames-1 {
			css.WriteString(fmt.Sprintf(
				"    .f%d { animation: show %.3fs ease forwards; animation-delay: %.3fs; opacity: 0; }\n",
				i, frameDuration, delay))
		} else {
			css.WriteString(fmt.Sprintf(
				"    .f%d { animation: hide %.3fs ease forwards; animation-delay: %.3fs; opacity: 0; }\n",
				i, frameDuration, delay))
		}
		texts.WriteString(fmt.Sprintf(
			`    <text class="f%d" x="%d" y="18" fill="#000009" text-anchor="middle" font-weight="800">%d days</text>`+"\n",
			i, valX, val))
	}

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
  <style>
    @keyframes show { from { opacity: 0; } to { opacity: 1; } }
    @keyframes hide { 0%% { opacity: 0; } 10%% { opacity: 1; } 90%% { opacity: 1; } 100%% { opacity: 0; } }
%s  </style>
  <rect width="%d" height="28" rx="6" fill="url(#bg)" filter="url(#shadow)"/>
  <rect x="%d" width="%d" height="28" rx="6" fill="url(#val)"/>
  <rect x="%d" width="6" height="28" fill="url(#val)"/>
  <g font-family="'Segoe UI',Roboto,Verdana,sans-serif" font-size="11" font-weight="600">
    <text x="%d" y="18" fill="#ffffff" text-anchor="middle">unemployed</text>
%s  </g>
</svg>`,
		totalWidth, valueText,
		css.String(),
		totalWidth,
		labelWidth, valueWidth,
		labelWidth,
		labelWidth/2,
		texts.String(),
	)

	c.Set("Content-Type", "image/svg+xml")
	c.Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=86400")
	return c.SendString(svg)
}
