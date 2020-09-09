package cynthia_test

import (
	"golang.org/x/tools/go/analysis"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/riita10069/cynthia"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	_t := testutil.Filter(t, func(format string, args ...interface{}) bool {
		for _, arg := range args {
			switch arg := arg.(type) {
			case *analysis.Pass:
				if !cynthia.IsTest(arg) {
					return false
				}
			}
		}
		return true
	})
	analysistest.Run(_t, testdata, cynthia.Analyzer, "a")
}
