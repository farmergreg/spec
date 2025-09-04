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
	packageName += ".Spec"
	code := fmt.Sprintf("%#v", a)
	r := code[len(packageName):]
	return r
}
