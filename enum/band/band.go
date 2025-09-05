package band

import "github.com/hamradiolog-net/adif-spec/v6/internal/codegen"

// Band represents a range of radio frequencies
type Band string

var _ codegen.CodeGenKey = Band("")

func (b Band) String() string {
	return string(b)
}
