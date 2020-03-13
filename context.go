package enrollment

// https://medium.com/@matryer/context-keys-in-go-5312346a868d for more details

// contextKey is our type for storing the Vehicle Entity in the context
type contextKey string

// String makes sure we are uniqe across the context
func (c contextKey) String() string {
	return "enrollment context key " + string(c)
}

var (
	contextKeyVehicleEntity = contextKey("vehicle-entity")
)