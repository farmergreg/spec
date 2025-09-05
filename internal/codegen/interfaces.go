package codegen

type CodeGenContainer interface {
	CodeGeneratorRecords() map[CodeGeneratorEnumValue]CodeGenSpec
	CodeGeneratorMetadata() CodeGeneratorMetadataForContainer
}

type CodeGeneratorMetadataForContainer struct {
	PackageName string
	DataType    string
}

type CodeGenSpec interface {
	CodeGeneratorMetadata() CodeGeneratorMetadataForEnum
}

type CodeGeneratorMetadataForEnum struct {
	ConstName     string
	ConstValue    string
	ConstComments string
	IsDeprecated  bool
}

type CodeGeneratorEnumValue interface {
	String() string
}
