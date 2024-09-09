package main

import (
	"os"
	"time"

	"whatever/m/15.maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
