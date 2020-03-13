package enrollment

import (
	"context"
	"errors"
)

// Enrollment encapsulates the whole enrollment process
type Enrollment interface {
	Enroll() Result
}

// Result signalize the result of Enrollment
// it can fail (e.g. user canceled) or succeed (optional returns)
type Result struct {
	Code          ResultCode
	ErrMsg        string
	VehicleEntity *VehicleEntity
}

// HandlerFunc type
type HandlerFunc func(context.Context) (context.Context, error)

// flows are static mapping between the flow types and necessary steps
// we need to perform
var flows = map[Flow][]HandlerFunc{
	FlowNoOp: {notImplementedHandler},
	FlowConnect: {
		connectPIC,
		connectAddVehicleToGarage,
		connectActivateVehicleByPIN,
	},
	FlowSmartlink: {
		smartlinkPIC,
		smartlinkAddVehicleToGarage,
		smartlinkActivateVehicle,
	},
	FlowSmartlinkConnect: {
		connectPIC,
		smartlinkConnectPIC,
		connectActivateVehicleByPIN,
		connectAddVehicleToGarage,
		smartlinkAddVehicleToGarage,
	},
}

// notImplementedHandler signalizes missing flow
func notImplementedHandler(ctx context.Context) (context.Context, error) {
	return ctx, errors.New("This flow is not implemented")
}
