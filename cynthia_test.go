package cynthia_test

import (
	"testing"

	"github.com/riita10069/cynthia"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, cynthia.Analyzer, "a")

}
