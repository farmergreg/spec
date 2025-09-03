package main

import (
	"fmt"
	"html/template"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"CommentsFor":    commentsFor,
	"Split":          strings.Split,
	"ToGoCode":       toGoCode,
	"ToGoIdentifier": toGoIdentifier,
}

// commentsFor retrieves comments for a given specification using the provided ViewBag.
func commentsFor(viewBag ViewBag, a any) string {
	return viewBag.Comments(a)
}

// toGoCode converts a struct to its Go code representation.
// it removes the package prefix from the output because the code will be used in the same package.
func toGoCode(a any, packageName string) string {
	packageName += ".Spec"
	code := fmt.Sprintf("%#v", a)
	r := code[len(packageName):]
	return r
}

// toGoIdentifier converts a string to a valid Go identifier by replacing or removing invalid characters.
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
