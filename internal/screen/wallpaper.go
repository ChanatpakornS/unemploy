package screen

import (
	"fmt"
	"strings"
)

func GenerateWallpaper(
	width int,
	height int,
	days int,
	cssFrames string,
	textFrames string,
) string {
	// Calculate responsive sizing based on dimensions
	fontSize := calculateFontSize(width, height)
	headerFontSize := fontSize / 5
	footerFontSize := fontSize / 5

	centerX := width / 2
	centerY := height / 2
	headerY := centerY - fontSize
	footerY := centerY + fontSize + headerFontSize + 20

	svg := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" role="img" aria-label="unemployed for %d days">
  <defs>
    <linearGradient id="bg" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0%%" stop-color="#0f172a"/>
      <stop offset="50%%" stop-color="#1e293b"/>
      <stop offset="100%%" stop-color="#334155"/>
    </linearGradient>
    <filter id="glow">
      <feGaussianBlur stdDeviation="4" result="coloredBlur"/>
      <feMerge>
        <feMergeNode in="coloredBlur"/>
        <feMergeNode in="SourceGraphic"/>
      </feMerge>
    </filter>
  </defs>
  <style>
    @keyframes show { from { opacity: 0; } to { opacity: 1; } }
    @keyframes hide { 0%% { opacity: 0; } 10%% { opacity: 1; } 90%% { opacity: 1; } 100%% { opacity: 0; } }
%s  </style>
  <rect width="%d" height="%d" fill="url(#bg)"/>
  <g font-family="'Segoe UI',Roboto,'Helvetica Neue',Arial,sans-serif">
    <text x="%d" y="%d" text-anchor="middle" font-size="%d" fill="#94a3b8" letter-spacing="2" font-weight="300">I'VE BEEN UNEMPLOYED FOR</text>
%s    <text x="%d" y="%d" text-anchor="middle" font-size="%d" fill="#64748b" letter-spacing="1" font-weight="300">DAYS</text>
  </g>
</svg>`,
		width, height, days,
		cssFrames,
		width, height,
		centerX, headerY, headerFontSize,
		textFrames,
		centerX, footerY, footerFontSize,
	)

	return svg
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

func generateWallpaperFrames(days int, width int, height int) (string, string) {
	fontSize := calculateFontSize(width, height)
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
		// Ease out exponential
		eased := 1.0 - (1.0-t)*(1.0-t)*(1.0-t)
		val := int(eased * float64(days))
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
