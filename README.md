# Go Merge SVG

<p align="center">
<b>Combine Multiple SVGS</b> <br> into one SVG element
</p>

A free utility to combine multiple SVGs into one SVG element, written in GO.


**Quick Setup:** One command to install lighweight binary.

**Blazing speeds:** Combines multiple SVGs in just 10 seconds.

**Direction:** Multiple options to direction horizontal, vertical available.

**Dependency Free:** No need to install any dependencies from `pip`, `npm`. Just download and run.

# C.I

![go-build-time](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=go-build-time&branch=master)
![go-test-runtime](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=go-test-runtime&branch=master)
![coverage](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=coverage&branch=master)
![go-binary-size](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=go-binary-size&branch=master)
![go-mod-dependencies](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=go-mod-dependencies&branch=master)
![allocs-per-op](https://coveritup.app/badge?org=kevincobain2000&repo=go-merge-svg&type=allocs-per-op&branch=master)

---

![go-build-time](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=go-build-time&output=svg&width=160&height=160&branch=master&line=fill)
![go-test-runtime](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=go-test-runtime&output=svg&width=160&height=160&branch=master&line=fill)
![coverage](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=coverage&output=svg&width=160&height=160&branch=master&line=fill)
![go-binary-size](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=go-binary-size&output=svg&width=160&height=160&branch=master&line=fill)
![go-mod-dependencies](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=go-mod-dependencies&output=svg&width=160&height=160&branch=master&line=fill)
![allocs-per-op](https://coveritup.app/chart?org=kevincobain2000&repo=go-merge-svg&type=allocs-per-op&output=svg&width=160&height=160&branch=master&line=fill)


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