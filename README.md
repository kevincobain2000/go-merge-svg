# Go Merge SVG

<p align="center">
<b>Combine Multiple SVGS</b> <br> into one SVG element
</p>

A free utility to combine multiple SVGs into one SVG element, written in GO.


**Quick Setup:** One command to install lighweight binary.

**Blazing speeds:** Combines multiple SVGs in just 10 seconds.

**Direction:** Multiple options to direction horizontal, vertical available.

**Dependecy Free:** No need to install any dependencies from `pip`, `npm`. Just download and run.

# C.I



# Install

```sh
curl -sL https://raw.githubusercontent.com/kevincobain2000/go-merge-svg/master/install.sh | sh
mv gms /usr/local/bin/gms
```


# CLI Usage

**Simple usage**

```sh
gms --files=svg1.svg --files=svg2.svg
```

**Advanced usages**


```sh
gms --files=svg1.svg --files=svg2.svg --direction=vertical --margin=10
```

**All Options**

```sh
  -direction string
    	direction of the merge (default "vertical")
  -files value
    	files to merge
  -margin float
    	margin between the SVGs
  -version
    	prints version
```

# Go Usage

Merge 2 svg strings

```go
import (
    "encoding/xml"
	svg "github.com/kevincobain2000/go-merge-svg/svg"
)

func main() {
    s := svg.NewSVG()
    c := s.MergeVertically("<svg xmlns=...", "<svg xmlns=...", 0)
    output, err := xml.MarshalIndent(c, "", "  ")
}
```

Merge more than 2 svg strings

```go
import (
    "encoding/xml"
	svg "github.com/kevincobain2000/go-merge-svg/svg"
)

func main() {
    s := svg.NewSVG()
    c1 := s.MergeVertically("<svg xmlns=...", "<svg xmlns=...", 0)
    c2 := s.MergeVertically(c1, "<svg xmlns=...", 0)
    c3 := s.MergeHorizontally(c2, "<svg xmlns=...", 0)
    output, err := xml.MarshalIndent(c3, "", "  ")
}
```

# CHANGE LOG

- v1.0 - initial release