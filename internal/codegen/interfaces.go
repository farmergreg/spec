package codegen

type CodeGenContainer interface {
	CodeGenMetadata() CodeGenContainerMetadata
	CodeGenRecords() map[CodeGenKey]CodeGenSpec
}

type CodeGenSpec interface {
	CodeGenMetadata() CodeGenEnumMetadata
}

type CodeGenKey interface {
	String() string
}

type CodeGenContainerMetadata struct {
	PackageName      string
	DataType         string
	CompositeKeyType string            // only used for enums that have composite keys
	CompositeKeyMap  map[string]string // key is the code, value is the comment; only used for enums that have composite keys
}

type CodeGenEnumMetadata struct {
	ConstName     string
	ConstValue    string
	ConstComments string
	IsDeprecated  bool
}
