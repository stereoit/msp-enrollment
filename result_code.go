package enrollment

// ResultCode signalizes success or fail of the whole Enrollment process
type ResultCode int

// Enrollment Code results
const (
	ResultCodeFail ResultCode = iota
	ResultCodeSuccess
)

// String returns human readable format
func (r ResultCode) String() string {
	return [...]string{"Fail", "Success"}[r]
}
