package screen

import "fmt"

func GenerateCard(
	days int,
	cssFrames string,
	textFrames string,
) string {
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

	return svg
}
