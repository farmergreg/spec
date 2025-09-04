package antpath

// AntPath represents the antenna path used for a QSO
type AntPath string

func (a AntPath) String() string {
	return string(a)
}
