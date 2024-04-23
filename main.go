package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	svg "github.com/kevincobain2000/go-merge-svg/svg"
)

var version = "dev"

type sliceFlags []string

func (i *sliceFlags) String() string {
	return "my string representation"
}

func (i *sliceFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Flags struct {
	version   bool
	direction string
	margin    float64
	cols      int
	rows      int
	files     sliceFlags
}

var f Flags

func main() {
	SetupFlags()
	if f.version {
		fmt.Println(version)
		return
	}

	out := Merge(f)
	fmt.Println(string(out))
}

func Merge(f Flags) []byte {
	s := svg.NewSVG()
	svgs := loadSVGs(s, f.files)
	var combined svg.SVG
	for i := range svgs {
		if i == len(svgs)-1 {
			break
		}

		switch f.direction {
		case "vertical":
			combined = s.MergeVertically(svgs[i], svgs[i+1], f.margin)
		case "horizontal":
			combined = s.MergeHorizontally(svgs[i], svgs[i+1], f.margin)
		case "grid":
			combined = s.MergeGrid(svgs[i], svgs[i+1], f.margin, f.cols, f.rows)
		default:
			// Optionally handle an unknown direction
		}

		svgs[i+1] = combined
	}
	output, err := xml.MarshalIndent(combined, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling combined SVG: %v\n", err)
		os.Exit(1)
	}
	return output
}

func loadSVGs(s svg.SVG, files []string) []svg.SVG {
	svgs := []svg.SVG{}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
		svg, err := s.ParseSVG(data)
		if err != nil {
			fmt.Printf("Error parsing SVG: %v\n", err)
			os.Exit(1)
		}
		svgs = append(svgs, svg)
	}
	return svgs
}

func SetupFlags() {
	flag.BoolVar(&f.version, "version", false, "prints version")
	flag.StringVar(&f.direction, "direction", "vertical", "direction of the merge - horizontal, vertical or grid")
	flag.Var(&f.files, "files", "svg files path to merge")
	flag.Float64Var(&f.margin, "margin", 10, "margin between the SVGs")

	flag.IntVar(&f.cols, "cols", 0, "number of columns (in case --direction=grid)")
	flag.IntVar(&f.rows, "rows", 0, "number of rows (in case --direction=grid)")

	flag.Parse()

	if len(f.files) < 2 {
		fmt.Println("You need to provide at least 2 files to merge")
		os.Exit(1)
	}
	if f.direction != "vertical" && f.direction != "horizontal" && f.direction != "grid" {
		fmt.Println("Direction must be either vertical, horizontal or grid")
		os.Exit(1)
	}
}
