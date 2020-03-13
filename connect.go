package enrollment

import (
	"context"
	"fmt"
)

// here we hold all connect building block implementations

// connectPIC component just calls webview in real life
func connectPIC(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Performin Connect PIC\n")
	return ctx, nil
}

// connectAddVehicleToGarage performs activation of the vehicle on the MBB/ODP
// we simulate the car was created OK and insert the vehicle entity for this
func connectAddVehicleToGarage(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Adding Connect vehicle to the garage (MBB/ODP) - very compliated\n")
	ve := &VehicleEntity{"some-unique-vehicle-code-for-connect"}
	ctx = context.WithValue(ctx, contextKeyVehicleEntity, ve)
	return ctx, nil
}

// connectActivateVehicleByPIN is the activation method for standard Connect car
func connectActivateVehicleByPIN(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Activating Connect vehicle by PIN\n")
	return ctx, nil
}

// connectActivateVehicleByPIN is the activation method for standard Connect car
func connectActivateVehicleByMileage(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Activating Connect vehicle by MILEAGE\n")
	return ctx, nil
}
