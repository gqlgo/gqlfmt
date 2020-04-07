package main

import (
	"bytes"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/parser"
	"golang.org/x/xerrors"
)

func process(filename string, src []byte, opt *Options) ([]byte, error) {
	source := &ast.Source{Name: filename, Input: string(src)}

	query, err := parser.ParseQuery(source)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return queryFormat(query), nil
}

func queryFormat(queryDocument *ast.QueryDocument) []byte {
	var buf bytes.Buffer
	astFormatter := formatter.NewFormatter(&buf)
	astFormatter.FormatQueryDocument(queryDocument)
	return buf.Bytes()
}