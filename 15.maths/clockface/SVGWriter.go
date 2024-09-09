package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

// SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	makeHand(w, secondHandPoint(t), secondHandLength, "#f00")
}

func minuteHand(w io.Writer, t time.Time) {
	makeHand(w, minuteHandPoint(t), minuteHandLength, "#000")
}

func hourHand(w io.Writer, t time.Time) {
	makeHand(w, hourHandPoint(t), hourHandLength, "#000")
}

func makeHand(w io.Writer, p Point, length float64, color string) {
	p = Point{p.X * length, p.Y * length}             // scale
	p = Point{p.X, -p.Y}                              // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} // translate
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`, p.X, p.Y, color)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
