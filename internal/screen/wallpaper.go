package screen

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func GenerateWallpaper(
	width int,
	height int,
	days int,
) ([]byte, error) {
	// Create a new context
	dc := gg.NewContext(width, height)

	// Draw gradient background
	drawGradientBackground(dc, width, height)

	// Load font
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	// Calculate responsive sizing
	fontSize := calculateFontSize(width, height)
	headerFontSize := float64(fontSize) / 5.0
	footerFontSize := float64(fontSize) / 5.0

	centerX := float64(width) / 2.0
	centerY := float64(height) / 2.0

	// Draw header text: "I'VE BEEN UNEMPLOYED FOR"
	headerY := centerY - float64(fontSize)
	face := truetype.NewFace(font, &truetype.Options{Size: headerFontSize})
	dc.SetFontFace(face)
	dc.SetColor(color.RGBA{148, 163, 184, 255}) // #94a3b8
	dc.DrawStringAnchored("I'VE BEEN UNEMPLOYED FOR", centerX, headerY, 0.5, 0.5)

	// Draw main day counter with glow effect
	daysY := centerY + float64(fontSize)/3.0
	daysFace := truetype.NewFace(font, &truetype.Options{Size: float64(fontSize)})
	dc.SetFontFace(daysFace)

	// Draw glow effect (multiple passes with decreasing opacity)
	daysText := fmt.Sprintf("%d", days)
	for i := range 8 {
		offset := float64(i) * 0.5
		alpha := uint8(20 - i*2)
		dc.SetColor(color.RGBA{241, 245, 249, alpha})
		dc.DrawStringAnchored(daysText, centerX+offset, daysY+offset, 0.5, 0.5)
		dc.DrawStringAnchored(daysText, centerX-offset, daysY+offset, 0.5, 0.5)
		dc.DrawStringAnchored(daysText, centerX+offset, daysY-offset, 0.5, 0.5)
		dc.DrawStringAnchored(daysText, centerX-offset, daysY-offset, 0.5, 0.5)
	}

	// Draw main text on top
	dc.SetColor(color.RGBA{241, 245, 249, 255}) // #f1f5f9
	dc.DrawStringAnchored(daysText, centerX, daysY, 0.5, 0.5)

	// Draw footer text: "DAYS"
	footerY := centerY + float64(fontSize) + headerFontSize + 20
	footerFace := truetype.NewFace(font, &truetype.Options{Size: footerFontSize})
	dc.SetFontFace(footerFace)
	dc.SetColor(color.RGBA{100, 116, 139, 255}) // #64748b
	dc.DrawStringAnchored("DAYS", centerX, footerY, 0.5, 0.5)

	// Encode to PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode PNG: %w", err)
	}

	return buf.Bytes(), nil
}

func drawGradientBackground(dc *gg.Context, width int, height int) {
	// Create gradient from top-left to bottom-right
	// Colors: #0f172a -> #1e293b -> #334155
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			// Calculate position in gradient (0.0 to 1.0)
			t := math.Sqrt(float64(x*x+y*y)) / math.Sqrt(float64(width*width+height*height))

			var r, g, b uint8
			if t < 0.5 {
				// Interpolate between #0f172a and #1e293b
				factor := t * 2.0
				r = uint8(15 + (30-15)*factor)
				g = uint8(23 + (41-23)*factor)
				b = uint8(42 + (59-42)*factor)
			} else {
				// Interpolate between #1e293b and #334155
				factor := (t - 0.5) * 2.0
				r = uint8(30 + (51-30)*factor)
				g = uint8(41 + (65-41)*factor)
				b = uint8(59 + (85-59)*factor)
			}

			img.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	dc.DrawImage(img, 0, 0)
}

func calculateFontSize(width int, height int) int {
	// Base calculation on the smaller dimension for better responsiveness
	smallerDim := width
	if height < width {
		smallerDim = height
	}

	// Calculate font size as a percentage of the smaller dimension
	fontSize := smallerDim / 8

	// Set reasonable bounds
	if fontSize < 48 {
		fontSize = 48
	}
	if fontSize > 200 {
		fontSize = 200
	}

	return fontSize
}
