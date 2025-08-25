package main

import (
	"html/template"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"Split":                 strings.Split,
	"ConvertToGoIdentifier": convertToGoIdentifier,
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
