package adifield

import "testing"

func TestField_IsValid(t *testing.T) {
	tests := []struct {
		name string
		f    Field
		want bool
	}{
		{name: "valid", f: QSO_DATE, want: true},
		{name: "invalid", f: "QSO_dATE", want: false},
		{name: "app", f: APP_LOTW_LASTQSL, want: true},
		{name: "userdef", f: "USERDEF1_QSO_DATE", want: true},
		{name: "bad", f: "this_is_not_a_field", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsValid(); got != tt.want {
				t.Errorf("Field.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
