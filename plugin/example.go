// This must be package main
package main

import (
	linters "github.com/austinvalle/go-internals-sandbox"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		linters.TagAnalyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
