package enrollment

import "testing"

func TestEnrollmentByVin(t *testing.T) {
	var tests = []struct {
		vin        string
		resultCode ResultCode
	}{
		{"TMBAJ-SMARTLINK", ResultCodeSuccess},
		{"TMBAJ-SMARTLINK_CONNECT", ResultCodeSuccess},
		{"TMBAJ-CONNECT", ResultCodeSuccess},
		{"TMBAJ-ECITIGO", ResultCodeFail},
	}

	for _, tt := range tests {
		enrollment := NewEnrollByVIN(tt.vin)
		result := enrollment.Enroll()
		if result.Code != tt.resultCode {
			t.Errorf("NewEnrollByVIN(\"%v\") expected: %v got: %v", tt.vin, tt.resultCode, result.Code)
		}
	}
}
