package main

import (
	"fmt"
	"html/template"
	"strings"
)

var tmplFuncs = template.FuncMap{
	"Split":    strings.Split,
	"ToGoCode": toGoCode,
}

func toGoCode(a any, packageName string) string {
	code := fmt.Sprintf("%#v", a)
	r := code[len(packageName+".Spec"):]

	if packageName != "mode" {
		// mode and submode don't get along because they both have the word "mode" in them.
		// We know that mode doesn't need this, so we skip it...
		r = strings.ReplaceAll(r, packageName+".", "")
	}

	return r
}
