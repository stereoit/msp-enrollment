package enrollment

// Flow represent given flow type (smartlink, mod1, connect...)
type Flow int

// Supported flow enum
const (
	FlowNoOp = iota
	FlowConnect
	FlowSmartlink
	FlowSmartlinkConnect
)

// String returns human redable version of given flow
func (f Flow) String() string {
	return [...]string{"Dummy", "Connect", "Smartlink", "Smartlink + Connect"}[f]
}
