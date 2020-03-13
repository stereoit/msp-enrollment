# Enrollment concept

This package showcase one possible solution to handle Enrollment wizards.


## Build

    $ make build


## Run

    $ msp-enroll <vin>

For example:

```bash
 $ ./msp-enroll TMBAJ-CONNECT            
Executing "Connect" flow
	- Performin Connect PIC
	- Adding Connect vehicle to the garage (MBB/ODP) - very compliated
	- Activating Connect vehicle by PIN
Vehicle "some-unique-vehicle-code-for-connect" VIN: "TMBAJ-CONNECT" was enrolled successfully.

 $ ./msp-enroll TMBAJ-SMARTLINK          
Executing "Smartlink" flow
	- Performin Smartlink Pic
	- Adding vehicle to the garage
	- Activating SL vehicle
Vehicle "some-unique-vehicle-code-for-smartlink" VIN: "TMBAJ-SMARTLINK" was enrolled successfully.

 $ ./msp-enroll TMBAJ-SMARTLINK_CONNECT  
Executing "Smartlink + Connect" flow
	- Performin Connect PIC
	- Custom SmartlinkPIC component that knows about Connect
	- Activating Connect vehicle by PIN
	- Adding Connect vehicle to the garage (MBB/ODP) - very compliated
	- Adding vehicle to the garage
Vehicle "some-unique-vehicle-code-for-smartlink" VIN: "TMBAJ-SMARTLINK_CONNECT" was enrolled successfully.

 $ ./msp-enroll TMBAJ-ECITIGO
Executing "Dummy" flow
Enrollment for VIN "TMBAJ-ECITIGO" failed: This flow is not implemented
```

## Test

    $ make test