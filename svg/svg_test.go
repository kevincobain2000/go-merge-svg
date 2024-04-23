package svg

import (
	"testing"
)

func TestNewSVG(t *testing.T) {
	svg := NewSVG()
	if svg.Xmlns != "http://www.w3.org/2000/svg" {
		t.Errorf("Expected xmlns to be http://www.w3.org/2000/svg, got %s", svg.Xmlns)
	}
	if svg.Version != "1.1" {
		t.Errorf("Expected version to be 1.1, got %s", svg.Version)
	}
}

func TestParseSVG(t *testing.T) {
	t.Run("valid SVG", func(t *testing.T) {
		raw := []byte(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="100" height="100"></svg>`)
		svg := SVG{}
		_, err := svg.ParseSVG(raw)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("invalid SVG", func(t *testing.T) {
		raw := []byte(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="100" height="100></svg>`)
		svg := SVG{}
		_, err := svg.ParseSVG(raw)
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}

func TestMergeVertically(t *testing.T) {
	t.Run("valid SVGs", func(t *testing.T) {
		s := NewSVG()
		svg1 := SVG{
			Width:   100,
			Height:  100,
			ViewBox: "0 0 100 100",
			Content: []byte("content1"),
		}
		svg2 := SVG{
			Width:   100,
			Height:  100,
			ViewBox: "0 0 100 100",
			Content: []byte("content2"),
		}
		combined := s.MergeVertically(svg1, svg2, 10)
		if combined.Width != 100 {
			t.Errorf("Expected width to be 100, got %f", combined.Width)
		}
		if combined.Height != 210 {
			t.Errorf("Expected height to be 210, got %f", combined.Height)
		}
		if combined.ViewBox != "0 0 100.000000 210.000000" {
			t.Errorf("Expected viewBox to be 0 0 100.000000 210.000000, got %s", combined.ViewBox)
		}
		if string(combined.Content) != "content1\n<g transform=\"translate(0, 110.000000)\">content2</g>" {
			t.Errorf("Expected content to be 'content1\n<g transform=\"translate(0, 110.000000)\">content2</g>', got %s", combined.Content)
		}
	})
}

func TestMergeHorizontally(t *testing.T) {
	t.Run("valid SVGs", func(t *testing.T) {
		s := NewSVG()
		svg1 := SVG{
			Width:   100,
			Height:  100,
			ViewBox: "0 0 100 100",
			Content: []byte("content1"),
		}
		svg2 := SVG{
			Width:   100,
			Height:  100,
			ViewBox: "0 0 100 100",
			Content: []byte("content2"),
		}
		combined := s.MergeHorizontally(svg1, svg2, 10)
		if combined.Width != 210 {
			t.Errorf("Expected width to be 210, got %f", combined.Width)
		}
		if combined.Height != 100 {
			t.Errorf("Expected height to be 100, got %f", combined.Height)
		}
		if combined.ViewBox != "0.000000 0.000000 210.000000 100.000000" {
			t.Errorf("Expected viewBox to be 0.000000 0.000000 210.000000 100.000000, got %s", combined.ViewBox)
		}
		if string(combined.Content) != "content1\n<g transform=\"translate(110.000000, 0)\">content2</g>" {
			t.Errorf("Expected content to be 'content1\n<g transform=\"translate(110.000000, 0)\">content2</g>', got %s", combined.Content)
		}
	})
}
