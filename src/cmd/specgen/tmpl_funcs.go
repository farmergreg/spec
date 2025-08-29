package main

import (
	"fmt"
	"html/template"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"Split":                 strings.Split,
	"ConvertToGoIdentifier": convertToGoIdentifier,
	"ConvertToGoCode":       convertToGoCode,
}

func convertToGoIdentifier(s string) string {
	var replacer = strings.NewReplacer(
		" ", "_",
		".", "_",
		"&", "AND",
		"(", "",
		")", "",
		",", "",
		"-", "_",
		"/", "_",
		"'", "",
		":", "_",
		"?", "Uncertain", // Special case for QSOComplete '?'
	)
	result := replacer.Replace(s)
	result = strings.ReplaceAll(result, "__", "_")
	result = strings.TrimSuffix(result, "_")
	return result
}

func convertToGoCode(a any) string {
	return fmt.Sprintf("%#v", a)
}
