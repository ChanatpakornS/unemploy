package screen

import (
	"fmt"
	"strings"
)

func GenerateBadge(
	totalWidth int,
	valueText string,
	css strings.Builder,
	labelWidth int,
	valueWidth int,
	texts strings.Builder,
) string {
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

	return svg
}
