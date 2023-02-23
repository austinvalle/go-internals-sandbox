package linters

import (
	"go/ast"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var TagAnalyzer = &analysis.Analyzer{
	Name:     "av-tag",
	Doc:      "av-tag checker",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.StructType)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		structType, ok := node.(*ast.StructType)
		if !ok {
			return
		}

		if structType.Fields == nil || structType.Fields.NumFields() < 1 {
			return
		}

		for _, field := range structType.Fields.List {
			if field.Tag == nil {
				continue
			}

			ok := hasTag(field.Tag, "av")
			if !ok {
				// skip when no struct tag for the key
				continue
			}

			paramType, ok := field.Type.(*ast.Ident)
			if !ok { // first param type isn't identificator so it can't be of type "string"
				return
			}

			if paramType.Name != "string" {
				pass.Reportf(field.Tag.Pos(), "cannot use '%s' tag on a non string type", "av")
			}
		}
	})
	return nil, nil
}

func hasTag(tag *ast.BasicLit, key string) bool {
	raw := strings.Trim(tag.Value, "`")

	_, ok := reflect.StructTag(raw).Lookup(key)

	return ok
}
