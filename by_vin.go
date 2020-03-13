package enrollment

import "strings"

// NewEnrollByVIN returns new instance of enrollment by vin implementation
// another implementation can by VehicleFromTheList (chosing virtual vehicle)
func NewEnrollByVIN(vin string) Enrollment {
	return &enrollmentByVIN{
		vin: vin,
	}
}

type enrollmentByVIN struct {
	vin string
}

// Enroll implements the Enrollment interface
// basically we first somehow identify the requested flow (in our example stubbed by `getFlowByVin`)
// then we execute the flow
// and return its Result
func (e *enrollmentByVIN) Enroll() Result {
	flow := getFlowByVin(e.vin)
	fh := NewFlowHandler(flow)
	return fh.Execute()
}

// getFlowbyVin is not clever, in real world we can use stuff like
// api.connect.skoda-auto.cz/api/v1/vehicle/<VIN> to check supported capabilities
// but for demonstration I guess this is OK
func getFlowByVin(vin string) Flow {
	if strings.Contains(vin, "SMARTLINK_CONNECT") {
		return FlowSmartlinkConnect
	}

	if strings.Contains(vin, "SMARTLINK") {
		return FlowSmartlink
	}

	if strings.Contains(vin, "CONNECT") {
		return FlowConnect
	}

	return FlowNoOp
}
