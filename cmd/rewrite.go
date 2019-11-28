package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/fatih/astrewrite"
	"github.com/sirupsen/logrus"
	"github.com/wwq1988/errors"
)

var imported bool
var importStr = `"github.com/wwq1988/errors"`

func rewrite(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Trace(err)
	}
	fset := token.NewFileSet()
	oldAST, err := parser.ParseFile(fset, filename, data, parser.ParseComments)
	if err != nil {
		return errors.TraceWithField(err, "filename", filename)
	}
	newAST := astrewrite.Walk(oldAST, visitor)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, newAST)
	if err != nil {
		return errors.Trace(err)
	}
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func visitor(n ast.Node) (ast.Node, bool) {
	switch v := n.(type) {
	case *ast.GenDecl:
		return handleImportDecl(v)
	case *ast.ReturnStmt:
		return handleReturn(v), true
	default:
		return n, true
	}
}

func handleImportDecl(gd *ast.GenDecl) (ast.Node, bool) {
	if gd.Tok != token.IMPORT {
		return gd, true
	}
	if imported {
		return gd, true
	}
	newSpecs := []ast.Spec{}
	found := false
	for _, s := range gd.Specs {
		im, ok := s.(*ast.ImportSpec)
		if !ok {
			continue
		}
		if im.Path.Value == `"errors"` {
			continue
		}
		if im.Path.Value == importStr {
			found = true
		}
		newSpecs = append(newSpecs, s)
	}
	if !found {
		newSpecs = append(newSpecs, &ast.ImportSpec{Path: &ast.BasicLit{Kind: token.STRING, Value: importStr}})
	}
	gd.Specs = newSpecs
	imported = true
	return gd, true
}

func handleReturn(rst *ast.ReturnStmt) ast.Node {
	for resultIdx, result := range rst.Results {
		ident, ok := result.(*ast.Ident)
		if !ok {
			continue
		}
		if ident.Obj == nil {
			continue
		}
		assignStmt, ok := ident.Obj.Decl.(*ast.AssignStmt)
		if !ok {
			continue
		}
		idx := checkError(assignStmt.Rhs)
		if idx == -1 {
			continue
		}
		rst.Results[resultIdx] = &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "errors"},
				Sel: &ast.Ident{Name: "Trace"},
			},
			Args: []ast.Expr{ident},
		}
		break
	}
	return rst
}

func checkError(rhs []ast.Expr) int {
	idx := -1
	found := false
outter:
	for _, rh := range rhs {
		callExpr, ok := rh.(*ast.CallExpr)
		if !ok {
			continue
		}
		fun := callExpr.Fun
		ident, ok := fun.(*ast.Ident)
		if !ok {
			continue
		}
		decl := ident.Obj.Decl
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		for subIdx, each := range funcDecl.Type.Results.List {
			ident, ok := each.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if ident.Name == "error" {
				idx += subIdx
				found = true
				break outter
			}
		}
		idx += len(funcDecl.Type.Results.List)
	}
	if found {
		idx += 1
	}
	return idx
}

func main() {
	if err := rewrite("t.go"); err != nil {
		logrus.Error(err)
	}
}
