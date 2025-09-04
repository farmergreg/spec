package codegen

type CodeGenContainer interface {
	CodeGeneratorRecords() map[CodeGeneratorEnumValue]CodeGenSpec
	CodeGeneratorMetadata() CodeGeneratorMetadataForContainer
}

type CodeGeneratorMetadataForContainer struct {
	PackageName string
}

type CodeGenSpec interface {
	CodeGeneratorMetadata() CodeGeneratorMetadataForEnum
}

type CodeGeneratorMetadataForEnum struct {
	ConstName     string
	ConstDataType string
	ConstValue    string
	Comments      string
}

type CodeGeneratorEnumValue interface {
	String() string
}
