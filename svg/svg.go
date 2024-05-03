package svg

import (
	"encoding/xml"
	"fmt"
	"strings"
	"sync"
)

// SVG represents the root of an SVG file.
type SVG struct {
	XMLName xml.Name
	Content []byte  `xml:",innerxml"`
	Width   float64 `xml:"width,attr"`
	Height  float64 `xml:"height,attr"`
	ViewBox string  `xml:"viewBox,attr"`
	Xmlns   string  `xml:"xmlns,attr"`
	Xlink   string  `xml:"xmlns:xlink,attr"`
	Version string  `xml:"version,attr"`
}

var svgPool = sync.Pool{
	New: func() interface{} {
		return &SVG{}
	},
}

// NewSVG creates a new SVG with default values.
func NewSVG() *SVG {
	svg := svgPool.Get().(*SVG)
	svg.XMLName = xml.Name{Local: "svg"}
	svg.Xmlns = "http://www.w3.org/2000/svg"
	svg.Xlink = "http://www.w3.org/1999/xlink"
	svg.Version = "1.1"
	return svg
}

// ParseSVG reads and parses an SVG file.
func (s *SVG) ParseSVG(raw []byte) (SVG, error) {
	var svg SVG
	err := xml.Unmarshal(raw, &svg)
	return svg, err
}

// MergeVertically combines two SVGs into one, arranging them vertically with a given margin.
func (s *SVG) MergeVertically(svg1, svg2 SVG, margin float64) SVG {
	combinedSVG := s

	// Calculate combined dimensions.
	combinedSVG.Width = svg1.Width // Assuming both SVGs have the same width for simplicity.
	combinedSVG.Height = svg1.Height + svg2.Height + margin

	// Assuming both SVGs have the same viewbox width for simplicity.
	viewBoxParts := strings.Fields(svg1.ViewBox)
	if len(viewBoxParts) == 4 {
		combinedSVG.ViewBox = fmt.Sprintf("%s %s %f %f", viewBoxParts[0], viewBoxParts[1], combinedSVG.Width, combinedSVG.Height)
	}

	// Combine the contents, adjusting the second SVG's position.
	//nolint: gocritic
	combinedSVG.Content = append(svg1.Content, []byte(fmt.Sprintf("\n<g transform=\"translate(0, %f)\">%s</g>", svg1.Height+margin, svg2.Content))...)

	return *combinedSVG
}

// MergeHorizontally combines two SVGs into one, arranging them horizontally with a given margin.
func (s *SVG) MergeHorizontally(svg1, svg2 SVG, margin float64) SVG {
	combinedSVG := s

	// Calculate combined dimensions.
	combinedSVG.Width = svg1.Width + svg2.Width + margin
	combinedSVG.Height = Max(svg1.Height, svg2.Height) // Take the taller SVG's height for the combined height.

	// Adjust the viewBox to accommodate both SVGs and the margin.
	viewBoxParts1 := strings.Fields(svg1.ViewBox)
	viewBoxParts2 := strings.Fields(svg2.ViewBox)
	if len(viewBoxParts1) == 4 && len(viewBoxParts2) == 4 {
		vbX1 := ParseFloat(viewBoxParts1[0])
		vbY1 := ParseFloat(viewBoxParts1[1])
		vbX2 := ParseFloat(viewBoxParts2[0])
		vbY2 := ParseFloat(viewBoxParts2[1])
		combinedVBX := Min(vbX1, vbX2)
		combinedVBY := Min(vbY1, vbY2)
		combinedSVG.ViewBox = fmt.Sprintf("%f %f %f %f", combinedVBX, combinedVBY, combinedSVG.Width, combinedSVG.Height)
	} else {
		// Fallback if viewBox is missing or malformed in either SVG.
		combinedSVG.ViewBox = fmt.Sprintf("0 0 %f %f", combinedSVG.Width, combinedSVG.Height)
	}

	// Combine the contents, adjusting the second SVG's position by translating it horizontally.
	//nolint: gocritic
	combinedSVG.Content = append(svg1.Content, []byte(fmt.Sprintf("\n<g transform=\"translate(%f, 0)\">%s</g>", svg1.Width+margin, svg2.Content))...)

	return *combinedSVG
}

// merge as grid
func (s *SVG) MergeGrid(svg1, svg2 SVG, margin float64, rows, cols int) SVG {
	combinedSVG := s

	// Calculate combined dimensions.
	combinedSVG.Width = svg1.Width*float64(cols) + margin*float64(cols-1)
	combinedSVG.Height = svg1.Height*float64(rows) + margin*float64(rows-1)

	// Adjust the viewBox to accommodate both SVGs and the margin.
	viewBoxParts1 := strings.Fields(svg1.ViewBox)
	viewBoxParts2 := strings.Fields(svg2.ViewBox)
	if len(viewBoxParts1) == 4 && len(viewBoxParts2) == 4 {
		vbX1 := ParseFloat(viewBoxParts1[0])
		vbY1 := ParseFloat(viewBoxParts1[1])
		vbX2 := ParseFloat(viewBoxParts2[0])
		vbY2 := ParseFloat(viewBoxParts2[1])
		combinedVBX := Min(vbX1, vbX2)
		combinedVBY := Min(vbY1, vbY2)
		combinedSVG.ViewBox = fmt.Sprintf("%f %f %f %f", combinedVBX, combinedVBY, combinedSVG.Width, combinedSVG.Height)
	} else {
		// Fallback if viewBox is missing or malformed in either SVG.
		combinedSVG.ViewBox = fmt.Sprintf("0 0 %f %f", combinedSVG.Width, combinedSVG.Height)
	}

	// Combine the contents, adjusting the second SVG's position by translating it horizontally.
	//nolint: gocritic
	combinedSVG.Content = append(svg1.Content, []byte(fmt.Sprintf("\n<g transform=\"translate(%f, 0)\">%s</g>", svg1.Width+margin, svg2.Content))...)

	return *combinedSVG
}
