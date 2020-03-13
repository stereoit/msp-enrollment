package enrollment

import (
	"context"
	"fmt"
)

// FlowHandler executes enrollment flow
type FlowHandler interface {
	Execute() Result
}

// NewFlowHandler returns instance of FlowHandler
// this is the wizard that steps through each blocks and safeguard the results
func NewFlowHandler(flow Flow) FlowHandler {
	return &flowHandler{
		flow,
	}
}

type flowHandler struct {
	flow Flow
}

// Execute implements FlowHandler interface
// this flow handler expects in the end there is a VehicleEntity presented in the
// context
func (f *flowHandler) Execute() Result {
	var err error
	ctx := context.Background()
	chain := flows[f.flow]
	fmt.Printf("Executing \"%v\" flow\n", f.flow)

	for _, step := range chain {
		ctx, err = step(ctx)

		if err != nil {
			return Result{
				Code:   ResultCodeFail,
				ErrMsg: err.Error(),
			}
		}
	}

	// read vehicle entity from context
	vehicleEntity, ok := ctx.Value(contextKeyVehicleEntity).(*VehicleEntity)
	if !ok {
		return Result{
			Code:   ResultCodeFail,
			ErrMsg: "Failed to extract Vehicle Entity from context",
		}
	}

	return Result{
		Code:          ResultCodeSuccess,
		VehicleEntity: vehicleEntity,
	}
}
