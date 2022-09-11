package main

import (
	"bytes"
	"fmt"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"
	"github.com/vektah/gqlparser/v2/parser"
)

func process(filename string, src []byte, opt *Options) ([]byte, error) {
	source := &ast.Source{Name: filename, Input: string(src)}

	query, err := parser.ParseQuery(source)
	if err == nil {
		return queryFormat(query), nil
	}

	schema, err := parser.ParseSchema(source)
	if err == nil {
		return schemaFormat(schema), nil
	}

	return nil, fmt.Errorf("%v is not GraphQL file: %w", filename, err)
}

func queryFormat(queryDocument *ast.QueryDocument) []byte {
	var buf bytes.Buffer
	astFormatter := formatter.NewFormatter(&buf)
	astFormatter.FormatQueryDocument(queryDocument)
	return buf.Bytes()
}

func schemaFormat(schemaDocument *ast.SchemaDocument) []byte {
	var buf bytes.Buffer
	astFormatter := formatter.NewFormatter(&buf)
	astFormatter.FormatSchemaDocument(schemaDocument)
	return buf.Bytes()
}
