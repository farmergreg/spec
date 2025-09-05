package credit

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// Credit represents an award credit identifier
type Credit string

var _ codegen.CodeGenKey = Credit("")

func (c Credit) String() string {
	return string(c)
}
