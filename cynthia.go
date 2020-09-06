package cynthia

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"strings"
)

const doc = "cynthia is a tool to check if the corresponding test exists"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "cynthia",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	functionFilter := []ast.Node{(*ast.FuncDecl)(nil)}

	// map[テスト関数の名前]bool
	signatureMap := make(map[string]bool)

	// テストの関数をhashに入れていく
	inspect.Preorder(functionFilter, func(testNode ast.Node) {
		switch testNode := testNode.(type) {
		case *ast.FuncDecl:
			testFuncObj := pass.TypesInfo.ObjectOf(testNode.Name)
			if testFuncObj == nil {break}
			if !testFuncObj.Exported() {break}
			if !strings.HasPrefix(testFuncObj.Name(), "Test") {break}

			signatureMap[testFuncObj.Name()] = true
		}
	})


	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:
			signatureObj := pass.TypesInfo.ObjectOf(funcNode.Name)
			if signatureObj == nil {break} // 後ろでぬるぽ踏まないように
			if !signatureObj.Exported() {break} // プライベートな関数はテストなくてもいいかな
			if strings.HasPrefix(signatureObj.Name(), "Test") {break} // Testのテストはいらない
			if !(signatureObj.Name() != "main" && signatureObj.Name() != "init") {break}
			matchTestName := fmt.Sprintf("Test%s", signatureObj.Name())

			if _, ok := signatureMap[matchTestName]; !ok {
				pass.Reportf(signatureObj.Pos(), "There is no test function implemented")
			}
		}
	})

	return nil, nil
}

