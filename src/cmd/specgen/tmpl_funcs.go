package main

import (
	"fmt"
	"html/template"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"Split":          strings.Split,
	"ToGoIdentifier": toGoIdentifier,
	"ToGoCode":       toGoCode,
}

func toGoIdentifier(s string) string {
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

func toGoCode(a any, packageName string) string {
	packageName += ".Spec"
	code := fmt.Sprintf("%#v", a)
	r := code[len(packageName):]
	return r
}
