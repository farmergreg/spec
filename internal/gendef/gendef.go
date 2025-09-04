package gendef

type GenDef interface {
	ConstName() string
	ConstValue() string
	Comments() string
}
