package unemploy

import (
	"fmt"
	"math"
	"strings"
	"time"
	"unemployed/internal/screen"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/unemploy", Unemploy)
	v1.Get("/unemploy/badge", UnemployBadge)
	v1.Get("/unemploy/wallpaper", UnemployWallpaper)
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

	svg := screen.GenerateCard(days, cssFrames, textFrames)

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

	svg := screen.GenerateBadge(totalWidth, valueText, css, labelWidth, valueWidth, texts)

	c.Set("Content-Type", "image/svg+xml")
	c.Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=86400")
	return c.SendString(svg)
}

func UnemployWallpaper(c fiber.Ctx) error {
	wallpaperRequest := new(WallpaperRequestParams)
	if err := c.Bind().Query(wallpaperRequest); err != nil {
		log.Error("Failed to parse wallpaper request", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message": "Invalid parameters",
		})
	}

	start, err := time.Parse("2006-01-02", wallpaperRequest.Start)
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

	cssFrames, textFrames := generateWallpaperFrames(days, wallpaperRequest.Width, wallpaperRequest.Height)

	svg := screen.GenerateWallpaper(
		wallpaperRequest.Width,
		wallpaperRequest.Height,
		days,
		cssFrames,
		textFrames,
	)

	c.Set("Content-Type", "image/svg+xml")
	c.Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=86400")
	return c.SendString(svg)
}

func generateWallpaperFrames(days int, width int, height int) (string, string) {
	fontSize := calculateWallpaperFontSize(width, height)
	centerX := width / 2
	centerY := height / 2
	daysY := centerY + fontSize/3

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
			`      <text class="f%d" x="%d" y="%d" text-anchor="middle" dominant-baseline="central" font-size="%d" font-weight="700" fill="#f1f5f9" filter="url(#glow)">%d</text>`+"\n",
			i, centerX, daysY, fontSize, val))
	}

	return css.String(), texts.String()
}

func calculateWallpaperFontSize(width int, height int) int {
	smallerDim := width
	if height < width {
		smallerDim = height
	}

	fontSize := smallerDim / 8

	if fontSize < 48 {
		fontSize = 48
	}
	if fontSize > 200 {
		fontSize = 200
	}

	return fontSize
}
