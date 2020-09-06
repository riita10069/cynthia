package main

import (
	"github.com/riita10069/cynthia"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(cynthia.Analyzer) }

