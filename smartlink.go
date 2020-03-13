package enrollment

import (
	"context"
	"fmt"
)

// here we hold all smartlink building block implementations

// smartlinkPIC make sure we go through smartlink user requirements
func smartlinkPIC(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Performin Smartlink Pic\n")
	return ctx, nil
}

// smartlinkAddVehicleToGarage performs the call to the SL garage
func smartlinkAddVehicleToGarage(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Adding vehicle to the garage\n")
	ve := &VehicleEntity{"some-unique-vehicle-code-for-smartlink"}
	ctx = context.WithValue(ctx, contextKeyVehicleEntity, ve)
	return ctx, nil
}

// smartlinkActivateVehicle send the Activate Patch reqquest to the GW
func smartlinkActivateVehicle(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Activating SL vehicle\n")
	return ctx, nil
}

// smartlinkConnectPIC knows is being called within Connect+Smarlink Flow and
// thus can handle things differently
func smartlinkConnectPIC(ctx context.Context) (context.Context, error) {
	fmt.Printf("\t- Custom SmartlinkPIC component that knows about Connect\n")
	return ctx, nil
}
