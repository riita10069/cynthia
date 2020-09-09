package cynthia

import (
	"errors"
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
	if !IsTest(pass) {
				return nil, errors.New("no test")
	}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	functionFilter := []ast.Node{(*ast.FuncDecl)(nil)}

	// map[テスト関数の名前]bool
	signatureMap := make(map[string]bool)

	// テストの関数をhashに入れていく
	inspect.Preorder(functionFilter, func(testNode ast.Node) {
		switch testNode := testNode.(type) {
		case *ast.FuncDecl:
			testFuncObj := pass.TypesInfo.ObjectOf(testNode.Name)
			if testFuncObj == nil {
				break
			}
			if !testFuncObj.Exported() {
				break
			}
			if !strings.HasPrefix(testFuncObj.Name(), "Test") {
				break
			}

			signatureMap[testFuncObj.Name()] = true
		}
	})

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:

			signatureObj := pass.TypesInfo.ObjectOf(funcNode.Name)
			if signatureObj == nil {
				break
			}
			if !signatureObj.Exported() {
				break
			}
			if strings.HasPrefix(signatureObj.Name(), "Test") {
				break
			}
			if strings.HasPrefix(signatureObj.Name(), "New") {
				break
			}
			if !(signatureObj.Name() != "main" && signatureObj.Name() != "init") {
				break
			}
			matchTestName := fmt.Sprintf("Test%s", signatureObj.Name())

			fmt.Println("match test name", matchTestName)
			fmt.Printf("---%v\n", signatureMap)
			if _, ok := signatureMap[matchTestName]; !ok {
				fmt.Println("falseの時", signatureObj.Pos())
				pass.Reportf(signatureObj.Pos(), "not implemented")
			}
		}
	})

	return nil, nil
}

func IsTest(pass *analysis.Pass) bool {
	if strings.HasSuffix(pass.Pkg.Path(), ".test") {
		return false
	}

	for _, f := range pass.Files {
		fn := pass.Fset.File(f.Pos()).Name()
		if strings.HasSuffix(fn, "_test.go") {
			return true
		}
	}

	return false
}
