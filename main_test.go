package main

import (
	"os"
	"strings"
	"testing"

	svg "github.com/kevincobain2000/go-merge-svg/svg"
)

func TestMerge(t *testing.T) {
	s := svg.NewSVG()
	t.Run("vertical", func(t *testing.T) {
		f := Flags{
			direction: "vertical",
			margin:    10,
			files:     []string{"testdata/1.svg", "testdata/2.svg"},
		}

		out := Merge(f, s)
		expected, err := os.ReadFile("testdata/1_2_vertical.svg")
		if err != nil {
			t.Fatal(err)
		}
		if minify(string(out)) != minify(string(expected)) {
			t.Errorf("Expected %s, got %s", expected, out)
		}
	})
	t.Run("horizontal", func(t *testing.T) {
		f := Flags{
			direction: "horizontal",
			margin:    10,
			files:     []string{"testdata/1.svg", "testdata/2.svg"},
		}

		out := Merge(f, s)
		expected, err := os.ReadFile("testdata/1_2_horizontal.svg")
		if err != nil {
			t.Fatal(err)
		}
		if minify(string(out)) != minify(string(expected)) {
			t.Errorf("Expected %s, got %s", expected, out)
		}
	})
}

func BenchmarkMerge(b *testing.B) {
	f := Flags{
		direction: "vertical",
		margin:    10,
		files:     []string{"testdata/1.svg", "testdata/2.svg"},
	}
	s := svg.NewSVG()

	for i := 0; i < b.N; i++ {
		Merge(f, s)
	}
}

func minify(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "\n", ""), " ", "")
}
